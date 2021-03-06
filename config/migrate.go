package config

import (
	"encoding/json"
	"fmt"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/migrasi/models"
	"github.com/opentracing/opentracing-go/log"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"os"
)

func (cfg *Config) InitMigrate() serror.SError {

	rethinkTable := os.Getenv("RETHINKDB_TABLE")
	cockroarchDb := os.Getenv("DB_NAME")

	query := fmt.Sprintf(`select coalesce(dv.driver_id::text, '')  as driver_id,
					   coalesce(dv.driver_name::text, '')  as driver_name,
					   coalesce(dv.driver_code::text, '')  as driver_code,
					   coalesce(dv.number_sim::text, '')   as number_sim,
					   coalesce(dv.phone_number::text, '') as phone_number,
					   coalesce(dv.member_id::text, '')    as member_id
				from %s.public.mt_driver dv`, cockroarchDb)

	rows, err := cfg.DB.Queryx(query)
	if err != nil {

	}

	defer rows.Close()
	for rows.Next() {
		data := models.CockroachModel{}
		if err := rows.StructScan(&data); err != nil {
			log.Error(serror.NewFromErrorc(err, fmt.Sprintf("Error StructScan %s", data)))
		}

		query := fmt.Sprintf(`select coalesce(mv.imei::text, '') as imei 
						from %s.public.mt_vehicle mv 
				  where mv.driver_id = $1`, cockroarchDb)

		err := cfg.DB.QueryRowx(query, data.DriverID).Scan(&data.Imei)
		if err != nil {
			log.Error(serror.NewFromErrorc(err, fmt.Sprintf("Error Scan %s", data.DriverID)))
		}

		fmt.Println(fmt.Sprintf("Replace vehicle_id %s and name %s rethinkdb", data.Imei, data.DriverName))

		result, err := r.Table(rethinkTable).Filter(map[string]interface{}{
			"vehicle_id": data.Imei,
			"name":       data.DriverName,
		}).Run(cfg.DBRething)

		res, err := result.Interface()
		rethingModel := []models.RethinkModel{}
		bodyBytes, _ := json.Marshal(res)
		err = json.Unmarshal(bodyBytes, &rethingModel)
		if err != nil {
			log.Error(serror.NewFromErrorc(err, fmt.Sprintf("Error Unmarshal %s", data.DriverID)))
		}

		if len(rethingModel) > 0 {
			_, err = r.Table(rethinkTable).Get(rethingModel[0].Id).Replace(map[string]interface{}{
				"id":               rethingModel[0].Id,
				"address":          "",
				"app_id":           "480cea73-3263-4111-bb7e-749b2c6493b3",
				"code":             data.DriverCode,
				"created_at":       r.Now(),
				"driving_status":   "-",
				"job_status":       "standby",
				"name":             data.DriverName,
				"owner_id":         data.MemberID,
				"phone":            data.PhoneNumber,
				"photo":            "",
				"sim_expired":      "",
				"sim_number":       data.NumberSim,
				"user_id":          "",
				"vehicle_id":       data.Imei,
				"runner_driver_id": data.DriverID,
			}).RunWrite(cfg.DBRething)
		} else {
			_, err = r.Table(rethinkTable).Insert(map[string]interface{}{
				"address":          "",
				"app_id":           "480cea73-3263-4111-bb7e-749b2c6493b3",
				"code":             data.DriverCode,
				"created_at":       r.Now(),
				"driving_status":   "-",
				"job_status":       "standby",
				"name":             data.DriverName,
				"owner_id":         data.MemberID,
				"phone":            data.PhoneNumber,
				"photo":            "",
				"sim_expired":      "",
				"sim_number":       data.NumberSim,
				"user_id":          "",
				"vehicle_id":       data.Imei,
				"runner_driver_id": data.DriverID,
			}).Run(cfg.DBRething)
		}

		if err != nil {
			log.Error(serror.NewFromErrorc(err, fmt.Sprintf("Error Input %s", data.DriverID)))
		}
	}

	return nil
}
