CREATE TABLE factory (
       factory_id			BIGSERIAL PRIMARY KEY,
       sprocket_production_actual 	INTEGER[] NOT NULL,
       sprocket_production_goal 	INTEGER[] NOT NULL,
       time 				INTEGER[] NOT NULL
);

CREATE TABLE sprocket (
       sprocket_id	BIGSERIAL PRIMARY KEY,
       teeth 		INTEGER NOT NULL,
       pitch_diameter 	INTEGER NOT NULL,
       outside_diameter INTEGER NOT NULL,
       pitch 		INTEGER NOT NULL
);
