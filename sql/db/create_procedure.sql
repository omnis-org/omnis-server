USE OMNIS;

DELIMITER // ;

-- get_perimeters
DROP PROCEDURE IF EXISTS get_perimeters//
CREATE PROCEDURE get_perimeters(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,description FROM Perimeter WHERE automatic=p_automatic;
END //


-- get_perimeter_by_id
DROP PROCEDURE IF EXISTS get_perimeter_by_id//
CREATE PROCEDURE get_perimeter_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,description FROM Perimeter WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_perimeters
DROP PROCEDURE IF EXISTS get_timestamp_perimeters//
CREATE PROCEDURE get_timestamp_perimeters(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  name_last_modification,description_last_modification
    FROM Perimeter
    WHERE automatic = p_automatic
    AND (name_last_modification IS NOT NULL OR description_last_modification IS NOT NULL);
END //


-- get_timestamp_perimeter_by_id
DROP PROCEDURE IF EXISTS get_timestamp_perimeter_by_id//
CREATE PROCEDURE get_timestamp_perimeter_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name_last_modification,description_last_modification
    FROM Perimeter
    WHERE automatic = p_automatic AND id = p_id
    AND (name_last_modification IS NOT NULL OR description_last_modification IS NOT NULL);
END //


-- insert_perimeter
DROP PROCEDURE IF EXISTS insert_perimeter//
CREATE PROCEDURE insert_perimeter(IN p_name VARCHAR(255),IN p_description TEXT, IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO Perimeter(automatic, name,description,  name_last_modification,description_last_modification)
    VALUES(p_automatic, p_name,p_description,
    CASE WHEN (p_name IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_description IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO Perimeter(id, automatic, name,description)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_name,p_description);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_perimeter_t
DROP TRIGGER IF EXISTS update_perimeter_t//
CREATE TRIGGER update_perimeter_t BEFORE UPDATE ON Perimeter
FOR EACH ROW BEGIN

IF NEW.name != OLD.name THEN SET NEW.name_last_modification = NOW(); END IF;

IF NEW.description != OLD.description THEN SET NEW.description_last_modification = NOW(); END IF;

END//


-- update_perimeter
DROP PROCEDURE IF EXISTS update_perimeter//
CREATE PROCEDURE update_perimeter(IN p_id INT, IN p_name VARCHAR(255),IN p_description TEXT, IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE Perimeter SET
        name = COALESCE(p_name, name),
        description = COALESCE(p_description, description)
        WHERE id = p_id AND automatic = true;

        UPDATE Perimeter SET
        -- name
        name = CASE WHEN (name_last_modification IS NULL) THEN p_name ELSE name END,
        name_last_modification = CASE WHEN (name_last_modification IS NULL) THEN NULL ELSE  name_last_modification END,
        -- description
        description = CASE WHEN (description_last_modification IS NULL) THEN p_description ELSE description END,
        description_last_modification = CASE WHEN (description_last_modification IS NULL) THEN NULL ELSE  description_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE Perimeter SET
        name = COALESCE(p_name, name),
        description = COALESCE(p_description, description)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_perimeter
DROP PROCEDURE IF EXISTS delete_perimeter//
CREATE PROCEDURE delete_perimeter(IN p_id INT)
BEGIN
    DELETE FROM Perimeter WHERE id=p_id;
END //

-- get_perimeter_by_name
DROP PROCEDURE IF EXISTS get_perimeter_by_name;
CREATE PROCEDURE get_perimeter_by_name(IN p_name VARCHAR(255), IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,description FROM Perimeter WHERE name=p_name AND automatic=p_automatic;
END //



-- get_locations
DROP PROCEDURE IF EXISTS get_locations//
CREATE PROCEDURE get_locations(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,description FROM Location WHERE automatic=p_automatic;
END //


-- get_location_by_id
DROP PROCEDURE IF EXISTS get_location_by_id//
CREATE PROCEDURE get_location_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,description FROM Location WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_locations
DROP PROCEDURE IF EXISTS get_timestamp_locations//
CREATE PROCEDURE get_timestamp_locations(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  name_last_modification,description_last_modification
    FROM Location
    WHERE automatic = p_automatic
    AND (name_last_modification IS NOT NULL OR description_last_modification IS NOT NULL);
END //


-- get_timestamp_location_by_id
DROP PROCEDURE IF EXISTS get_timestamp_location_by_id//
CREATE PROCEDURE get_timestamp_location_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name_last_modification,description_last_modification
    FROM Location
    WHERE automatic = p_automatic AND id = p_id
    AND (name_last_modification IS NOT NULL OR description_last_modification IS NOT NULL);
END //


-- insert_location
DROP PROCEDURE IF EXISTS insert_location//
CREATE PROCEDURE insert_location(IN p_name VARCHAR(255),IN p_description TEXT, IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO Location(automatic, name,description,  name_last_modification,description_last_modification)
    VALUES(p_automatic, p_name,p_description,
    CASE WHEN (p_name IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_description IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO Location(id, automatic, name,description)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_name,p_description);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_location_t
DROP TRIGGER IF EXISTS update_location_t//
CREATE TRIGGER update_location_t BEFORE UPDATE ON Location
FOR EACH ROW BEGIN

IF NEW.name != OLD.name THEN SET NEW.name_last_modification = NOW(); END IF;

IF NEW.description != OLD.description THEN SET NEW.description_last_modification = NOW(); END IF;

END//


-- update_location
DROP PROCEDURE IF EXISTS update_location//
CREATE PROCEDURE update_location(IN p_id INT, IN p_name VARCHAR(255),IN p_description TEXT, IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE Location SET
        name = COALESCE(p_name, name),
        description = COALESCE(p_description, description)
        WHERE id = p_id AND automatic = true;

        UPDATE Location SET
        -- name
        name = CASE WHEN (name_last_modification IS NULL) THEN p_name ELSE name END,
        name_last_modification = CASE WHEN (name_last_modification IS NULL) THEN NULL ELSE  name_last_modification END,
        -- description
        description = CASE WHEN (description_last_modification IS NULL) THEN p_description ELSE description END,
        description_last_modification = CASE WHEN (description_last_modification IS NULL) THEN NULL ELSE  description_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE Location SET
        name = COALESCE(p_name, name),
        description = COALESCE(p_description, description)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_location
DROP PROCEDURE IF EXISTS delete_location//
CREATE PROCEDURE delete_location(IN p_id INT)
BEGIN
    DELETE FROM Location WHERE id=p_id;
END //

-- get_location_by_name
DROP PROCEDURE IF EXISTS get_location_by_name;
CREATE PROCEDURE get_location_by_name(IN p_name VARCHAR(255), IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,description FROM Location WHERE name=p_name AND automatic=p_automatic;
END //



-- get_operating_systems
DROP PROCEDURE IF EXISTS get_operating_systems//
CREATE PROCEDURE get_operating_systems(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,platform,platform_family,platform_version,kernel_version FROM OperatingSystem WHERE automatic=p_automatic;
END //


-- get_operating_system_by_id
DROP PROCEDURE IF EXISTS get_operating_system_by_id//
CREATE PROCEDURE get_operating_system_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,platform,platform_family,platform_version,kernel_version FROM OperatingSystem WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_operating_systems
DROP PROCEDURE IF EXISTS get_timestamp_operating_systems//
CREATE PROCEDURE get_timestamp_operating_systems(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  name_last_modification,platform_last_modification,platform_family_last_modification,platform_version_last_modification,kernel_version_last_modification
    FROM OperatingSystem
    WHERE automatic = p_automatic
    AND (name_last_modification IS NOT NULL OR platform_last_modification IS NOT NULL OR platform_family_last_modification IS NOT NULL OR platform_version_last_modification IS NOT NULL OR kernel_version_last_modification IS NOT NULL);
END //


-- get_timestamp_operating_system_by_id
DROP PROCEDURE IF EXISTS get_timestamp_operating_system_by_id//
CREATE PROCEDURE get_timestamp_operating_system_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name_last_modification,platform_last_modification,platform_family_last_modification,platform_version_last_modification,kernel_version_last_modification
    FROM OperatingSystem
    WHERE automatic = p_automatic AND id = p_id
    AND (name_last_modification IS NOT NULL OR platform_last_modification IS NOT NULL OR platform_family_last_modification IS NOT NULL OR platform_version_last_modification IS NOT NULL OR kernel_version_last_modification IS NOT NULL);
END //


-- insert_operating_system
DROP PROCEDURE IF EXISTS insert_operating_system//
CREATE PROCEDURE insert_operating_system(IN p_name VARCHAR(100),IN p_platform VARCHAR(100),IN p_platform_family VARCHAR(100),IN p_platform_version VARCHAR(100),IN p_kernel_version VARCHAR(100), IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO OperatingSystem(automatic, name,platform,platform_family,platform_version,kernel_version,  name_last_modification,platform_last_modification,platform_family_last_modification,platform_version_last_modification,kernel_version_last_modification)
    VALUES(p_automatic, p_name,p_platform,p_platform_family,p_platform_version,p_kernel_version,
    CASE WHEN (p_name IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_platform IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_platform_family IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_platform_version IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_kernel_version IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO OperatingSystem(id, automatic, name,platform,platform_family,platform_version,kernel_version)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_name,p_platform,p_platform_family,p_platform_version,p_kernel_version);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_operating_system_t
DROP TRIGGER IF EXISTS update_operating_system_t//
CREATE TRIGGER update_operating_system_t BEFORE UPDATE ON OperatingSystem
FOR EACH ROW BEGIN

IF NEW.name != OLD.name THEN SET NEW.name_last_modification = NOW(); END IF;

IF NEW.platform != OLD.platform THEN SET NEW.platform_last_modification = NOW(); END IF;

IF NEW.platform_family != OLD.platform_family THEN SET NEW.platform_family_last_modification = NOW(); END IF;

IF NEW.platform_version != OLD.platform_version THEN SET NEW.platform_version_last_modification = NOW(); END IF;

IF NEW.kernel_version != OLD.kernel_version THEN SET NEW.kernel_version_last_modification = NOW(); END IF;

END//


-- update_operating_system
DROP PROCEDURE IF EXISTS update_operating_system//
CREATE PROCEDURE update_operating_system(IN p_id INT, IN p_name VARCHAR(100),IN p_platform VARCHAR(100),IN p_platform_family VARCHAR(100),IN p_platform_version VARCHAR(100),IN p_kernel_version VARCHAR(100), IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE OperatingSystem SET
        name = COALESCE(p_name, name),
        platform = COALESCE(p_platform, platform),
        platform_family = COALESCE(p_platform_family, platform_family),
        platform_version = COALESCE(p_platform_version, platform_version),
        kernel_version = COALESCE(p_kernel_version, kernel_version)
        WHERE id = p_id AND automatic = true;

        UPDATE OperatingSystem SET
        -- name
        name = CASE WHEN (name_last_modification IS NULL) THEN p_name ELSE name END,
        name_last_modification = CASE WHEN (name_last_modification IS NULL) THEN NULL ELSE  name_last_modification END,
        -- platform
        platform = CASE WHEN (platform_last_modification IS NULL) THEN p_platform ELSE platform END,
        platform_last_modification = CASE WHEN (platform_last_modification IS NULL) THEN NULL ELSE  platform_last_modification END,
        -- platform_family
        platform_family = CASE WHEN (platform_family_last_modification IS NULL) THEN p_platform_family ELSE platform_family END,
        platform_family_last_modification = CASE WHEN (platform_family_last_modification IS NULL) THEN NULL ELSE  platform_family_last_modification END,
        -- platform_version
        platform_version = CASE WHEN (platform_version_last_modification IS NULL) THEN p_platform_version ELSE platform_version END,
        platform_version_last_modification = CASE WHEN (platform_version_last_modification IS NULL) THEN NULL ELSE  platform_version_last_modification END,
        -- kernel_version
        kernel_version = CASE WHEN (kernel_version_last_modification IS NULL) THEN p_kernel_version ELSE kernel_version END,
        kernel_version_last_modification = CASE WHEN (kernel_version_last_modification IS NULL) THEN NULL ELSE  kernel_version_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE OperatingSystem SET
        name = COALESCE(p_name, name),
        platform = COALESCE(p_platform, platform),
        platform_family = COALESCE(p_platform_family, platform_family),
        platform_version = COALESCE(p_platform_version, platform_version),
        kernel_version = COALESCE(p_kernel_version, kernel_version)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_operating_system
DROP PROCEDURE IF EXISTS delete_operating_system//
CREATE PROCEDURE delete_operating_system(IN p_id INT)
BEGIN
    DELETE FROM OperatingSystem WHERE id=p_id;
END //

-- get_operating_systems_by_name
DROP PROCEDURE IF EXISTS get_operating_systems_by_name;
CREATE PROCEDURE get_operating_systems_by_name(IN p_name VARCHAR(255), IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,platform,platform_family,platform_version,kernel_version FROM OperatingSystem WHERE name=p_name AND automatic=p_automatic;
END //



-- get_tags
DROP PROCEDURE IF EXISTS get_tags//
CREATE PROCEDURE get_tags(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,color FROM Tag WHERE automatic=p_automatic;
END //


-- get_tag_by_id
DROP PROCEDURE IF EXISTS get_tag_by_id//
CREATE PROCEDURE get_tag_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,color FROM Tag WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_tags
DROP PROCEDURE IF EXISTS get_timestamp_tags//
CREATE PROCEDURE get_timestamp_tags(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  name_last_modification,color_last_modification
    FROM Tag
    WHERE automatic = p_automatic
    AND (name_last_modification IS NOT NULL OR color_last_modification IS NOT NULL);
END //


-- get_timestamp_tag_by_id
DROP PROCEDURE IF EXISTS get_timestamp_tag_by_id//
CREATE PROCEDURE get_timestamp_tag_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name_last_modification,color_last_modification
    FROM Tag
    WHERE automatic = p_automatic AND id = p_id
    AND (name_last_modification IS NOT NULL OR color_last_modification IS NOT NULL);
END //


-- insert_tag
DROP PROCEDURE IF EXISTS insert_tag//
CREATE PROCEDURE insert_tag(IN p_name VARCHAR(255),IN p_color VARCHAR(10), IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO Tag(automatic, name,color,  name_last_modification,color_last_modification)
    VALUES(p_automatic, p_name,p_color,
    CASE WHEN (p_name IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_color IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO Tag(id, automatic, name,color)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_name,p_color);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_tag_t
DROP TRIGGER IF EXISTS update_tag_t//
CREATE TRIGGER update_tag_t BEFORE UPDATE ON Tag
FOR EACH ROW BEGIN

IF NEW.name != OLD.name THEN SET NEW.name_last_modification = NOW(); END IF;

IF NEW.color != OLD.color THEN SET NEW.color_last_modification = NOW(); END IF;

END//


-- update_tag
DROP PROCEDURE IF EXISTS update_tag//
CREATE PROCEDURE update_tag(IN p_id INT, IN p_name VARCHAR(255),IN p_color VARCHAR(10), IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE Tag SET
        name = COALESCE(p_name, name),
        color = COALESCE(p_color, color)
        WHERE id = p_id AND automatic = true;

        UPDATE Tag SET
        -- name
        name = CASE WHEN (name_last_modification IS NULL) THEN p_name ELSE name END,
        name_last_modification = CASE WHEN (name_last_modification IS NULL) THEN NULL ELSE  name_last_modification END,
        -- color
        color = CASE WHEN (color_last_modification IS NULL) THEN p_color ELSE color END,
        color_last_modification = CASE WHEN (color_last_modification IS NULL) THEN NULL ELSE  color_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE Tag SET
        name = COALESCE(p_name, name),
        color = COALESCE(p_color, color)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_tag
DROP PROCEDURE IF EXISTS delete_tag//
CREATE PROCEDURE delete_tag(IN p_id INT)
BEGIN
    DELETE FROM Tag WHERE id=p_id;
END //


-- get_softwares
DROP PROCEDURE IF EXISTS get_softwares//
CREATE PROCEDURE get_softwares(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,version,is_intern FROM Software WHERE automatic=p_automatic;
END //


-- get_software_by_id
DROP PROCEDURE IF EXISTS get_software_by_id//
CREATE PROCEDURE get_software_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,version,is_intern FROM Software WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_softwares
DROP PROCEDURE IF EXISTS get_timestamp_softwares//
CREATE PROCEDURE get_timestamp_softwares(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  name_last_modification,version_last_modification,is_intern_last_modification
    FROM Software
    WHERE automatic = p_automatic
    AND (name_last_modification IS NOT NULL OR version_last_modification IS NOT NULL OR is_intern_last_modification IS NOT NULL);
END //


-- get_timestamp_software_by_id
DROP PROCEDURE IF EXISTS get_timestamp_software_by_id//
CREATE PROCEDURE get_timestamp_software_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name_last_modification,version_last_modification,is_intern_last_modification
    FROM Software
    WHERE automatic = p_automatic AND id = p_id
    AND (name_last_modification IS NOT NULL OR version_last_modification IS NOT NULL OR is_intern_last_modification IS NOT NULL);
END //


-- insert_software
DROP PROCEDURE IF EXISTS insert_software//
CREATE PROCEDURE insert_software(IN p_name VARCHAR(255),IN p_version VARCHAR(255),IN p_is_intern BOOLEAN, IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO Software(automatic, name,version,is_intern,  name_last_modification,version_last_modification,is_intern_last_modification)
    VALUES(p_automatic, p_name,p_version,p_is_intern,
    CASE WHEN (p_name IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_version IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_is_intern IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO Software(id, automatic, name,version,is_intern)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_name,p_version,p_is_intern);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_software_t
DROP TRIGGER IF EXISTS update_software_t//
CREATE TRIGGER update_software_t BEFORE UPDATE ON Software
FOR EACH ROW BEGIN

IF NEW.name != OLD.name THEN SET NEW.name_last_modification = NOW(); END IF;

IF NEW.version != OLD.version THEN SET NEW.version_last_modification = NOW(); END IF;

IF NEW.is_intern != OLD.is_intern THEN SET NEW.is_intern_last_modification = NOW(); END IF;

END//


-- update_software
DROP PROCEDURE IF EXISTS update_software//
CREATE PROCEDURE update_software(IN p_id INT, IN p_name VARCHAR(255),IN p_version VARCHAR(255),IN p_is_intern BOOLEAN, IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE Software SET
        name = COALESCE(p_name, name),
        version = COALESCE(p_version, version),
        is_intern = COALESCE(p_is_intern, is_intern)
        WHERE id = p_id AND automatic = true;

        UPDATE Software SET
        -- name
        name = CASE WHEN (name_last_modification IS NULL) THEN p_name ELSE name END,
        name_last_modification = CASE WHEN (name_last_modification IS NULL) THEN NULL ELSE  name_last_modification END,
        -- version
        version = CASE WHEN (version_last_modification IS NULL) THEN p_version ELSE version END,
        version_last_modification = CASE WHEN (version_last_modification IS NULL) THEN NULL ELSE  version_last_modification END,
        -- is_intern
        is_intern = CASE WHEN (is_intern_last_modification IS NULL) THEN p_is_intern ELSE is_intern END,
        is_intern_last_modification = CASE WHEN (is_intern_last_modification IS NULL) THEN NULL ELSE  is_intern_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE Software SET
        name = COALESCE(p_name, name),
        version = COALESCE(p_version, version),
        is_intern = COALESCE(p_is_intern, is_intern)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_software
DROP PROCEDURE IF EXISTS delete_software//
CREATE PROCEDURE delete_software(IN p_id INT)
BEGIN
    DELETE FROM Software WHERE id=p_id;
END //


-- get_machines
DROP PROCEDURE IF EXISTS get_machines//
CREATE PROCEDURE get_machines(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,uuid,hostname,label,description,virtualization_system,serial_number,machine_type,perimeter_id,location_id,operating_system_id,omnis_version FROM Machine WHERE automatic=p_automatic AND authorized=1;
END //


-- get_pending_machines
DROP PROCEDURE IF EXISTS get_pending_machines//
CREATE PROCEDURE get_pending_machines()
BEGIN
    SELECT id,uuid,authorized,hostname,label,description,virtualization_system,serial_number,machine_type,perimeter_id,location_id,operating_system_id,omnis_version FROM Machine WHERE authorized IS NULL AND automatic=0;
END //


-- get_machine_by_id
DROP PROCEDURE IF EXISTS get_machine_by_id//
CREATE PROCEDURE get_machine_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id,uuid,hostname,label,description,virtualization_system,serial_number,machine_type,perimeter_id,location_id,operating_system_id,omnis_version FROM Machine WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_machines
DROP PROCEDURE IF EXISTS get_timestamp_machines//
CREATE PROCEDURE get_timestamp_machines(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, uuid_last_modification, hostname_last_modification,label_last_modification,description_last_modification,virtualization_system_last_modification,serial_number_last_modification,machine_type_last_modification,perimeter_id_last_modification,location_id_last_modification,operating_system_id_last_modification,omnis_version_last_modification
    FROM Machine
    WHERE automatic = p_automatic
    AND (uuid_last_modification IS NOT NULL OR hostname_last_modification IS NOT NULL OR label_last_modification IS NOT NULL OR description_last_modification IS NOT NULL OR virtualization_system_last_modification IS NOT NULL OR serial_number_last_modification IS NOT NULL OR perimeter_id_last_modification IS NOT NULL OR location_id_last_modification IS NOT NULL OR operating_system_id_last_modification IS NOT NULL OR machine_type_last_modification IS NOT NULL OR omnis_version_last_modification IS NOT NULL);
END //


-- get_timestamp_machine_by_id
DROP PROCEDURE IF EXISTS get_timestamp_machine_by_id//
CREATE PROCEDURE get_timestamp_machine_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, uuid_last_modification, hostname_last_modification,label_last_modification,description_last_modification,virtualization_system_last_modification,serial_number_last_modification,machine_type_last_modification,perimeter_id_last_modification,location_id_last_modification,operating_system_id_last_modification,omnis_version_last_modification
    FROM Machine
    WHERE automatic = p_automatic AND id = p_id
    AND (uuid_last_modification IS NOT NULL OR hostname_last_modification IS NOT NULL OR label_last_modification IS NOT NULL OR description_last_modification IS NOT NULL OR virtualization_system_last_modification IS NOT NULL OR serial_number_last_modification IS NOT NULL OR perimeter_id_last_modification IS NOT NULL OR location_id_last_modification IS NOT NULL OR operating_system_id_last_modification IS NOT NULL OR machine_type_last_modification IS NOT NULL OR omnis_version_last_modification IS NOT NULL);
END //


-- insert_machine
DROP PROCEDURE IF EXISTS insert_machine//
CREATE PROCEDURE insert_machine(IN p_uuid VARCHAR(36), IN p_authorized BOOLEAN, IN p_hostname VARCHAR(255),IN p_label VARCHAR(255),IN p_description TEXT,IN p_virtualization_system VARCHAR(255),IN p_serial_number VARCHAR(255),IN p_machine_type VARCHAR(255),IN p_perimeter_id INT,IN p_location_id INT,IN p_operating_system_id INT,IN p_omnis_version VARCHAR(255), IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO Machine(automatic,uuid,authorized,hostname,label,description,virtualization_system,serial_number,machine_type,perimeter_id,location_id,operating_system_id,omnis_version, uuid_last_modification, authorized_last_modification, hostname_last_modification,label_last_modification,description_last_modification,virtualization_system_last_modification,serial_number_last_modification,machine_type_last_modification,perimeter_id_last_modification,location_id_last_modification,operating_system_id_last_modification,omnis_version_last_modification)
    VALUES(p_automatic, p_uuid, p_authorized, p_hostname,p_label,p_description,p_virtualization_system,p_serial_number,p_machine_type,p_perimeter_id,p_location_id,p_operating_system_id,p_omnis_version,
    CASE WHEN (p_uuid IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_authorized IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_hostname IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_label IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_description IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_virtualization_system IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_serial_number IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_machine_type IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_perimeter_id IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_location_id IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_operating_system_id IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_omnis_version IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO Machine(id, automatic, uuid, authorized, hostname,label,description,virtualization_system,serial_number,perimeter_id,location_id,operating_system_id,machine_type,omnis_version)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_uuid, p_authorized, p_hostname,p_label,p_description,p_virtualization_system,p_serial_number,p_perimeter_id,p_location_id,p_operating_system_id,p_machine_type,p_omnis_version);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_machine_t
DROP TRIGGER IF EXISTS update_machine_t//
CREATE TRIGGER update_machine_t BEFORE UPDATE ON Machine
FOR EACH ROW BEGIN
IF NEW.uuid != OLD.uuid THEN SET NEW.uuid_last_modification = NOW(); END IF;

IF NEW.authorized != OLD.authorized THEN SET NEW.authorized_last_modification = NOW(); END IF;

IF NEW.hostname != OLD.hostname THEN SET NEW.hostname_last_modification = NOW(); END IF;

IF NEW.label != OLD.label THEN SET NEW.label_last_modification = NOW(); END IF;

IF NEW.description != OLD.description THEN SET NEW.description_last_modification = NOW(); END IF;

IF NEW.virtualization_system != OLD.virtualization_system THEN SET NEW.virtualization_system_last_modification = NOW(); END IF;

IF NEW.serial_number != OLD.serial_number THEN SET NEW.serial_number_last_modification = NOW(); END IF;

IF NEW.machine_type != OLD.machine_type THEN SET NEW.machine_type_last_modification = NOW(); END IF;

IF NEW.perimeter_id != OLD.perimeter_id THEN SET NEW.perimeter_id_last_modification = NOW(); END IF;

IF NEW.location_id != OLD.location_id THEN SET NEW.location_id_last_modification = NOW(); END IF;

IF NEW.operating_system_id != OLD.operating_system_id THEN SET NEW.operating_system_id_last_modification = NOW(); END IF;

IF NEW.omnis_version != OLD.omnis_version THEN SET NEW.omnis_version_last_modification = NOW(); END IF;

END//


-- update_machine
DROP PROCEDURE IF EXISTS update_machine//
CREATE PROCEDURE update_machine(IN p_id INT, IN p_uuid VARCHAR(36), IN p_authorized BOOLEAN, IN p_hostname VARCHAR(255),IN p_label VARCHAR(255),IN p_description TEXT,IN p_virtualization_system VARCHAR(255),IN p_serial_number VARCHAR(255),IN p_machine_type VARCHAR(255),IN p_perimeter_id INT,IN p_location_id INT,IN p_operating_system_id INT,IN p_omnis_version VARCHAR(255), IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE Machine SET
        uuid = COALESCE(p_uuid, uuid),
        authorized = COALESCE(p_authorized, authorized),
        hostname = COALESCE(p_hostname, hostname),
        label = COALESCE(p_label, label),
        description = COALESCE(p_description, description),
        virtualization_system = COALESCE(p_virtualization_system, virtualization_system),
        serial_number = COALESCE(p_serial_number, serial_number),
        machine_type = COALESCE(p_machine_type, machine_type),
        perimeter_id = COALESCE(p_perimeter_id, perimeter_id),
        location_id = COALESCE(p_location_id, location_id),
        operating_system_id = COALESCE(p_operating_system_id, operating_system_id),
        omnis_version = COALESCE(p_omnis_version, omnis_version)
        WHERE id = p_id AND automatic = true;

        UPDATE Machine SET
        -- uuid
        uuid = CASE WHEN (uuid_last_modification IS NULL) THEN p_uuid ELSE uuid END,
        uuid_last_modification = CASE WHEN (uuid_last_modification IS NULL) THEN NULL ELSE  uuid_last_modification END,
        -- authorized
        authorized = CASE WHEN (authorized_last_modification IS NULL) THEN p_authorized ELSE authorized END,
        authorized_last_modification = CASE WHEN (authorized_last_modification IS NULL) THEN NULL ELSE  authorized_last_modification END,
        -- hostname
        hostname = CASE WHEN (hostname_last_modification IS NULL) THEN p_hostname ELSE hostname END,
        hostname_last_modification = CASE WHEN (hostname_last_modification IS NULL) THEN NULL ELSE  hostname_last_modification END,
        -- label
        label = CASE WHEN (label_last_modification IS NULL) THEN p_label ELSE label END,
        label_last_modification = CASE WHEN (label_last_modification IS NULL) THEN NULL ELSE  label_last_modification END,
        -- description
        description = CASE WHEN (description_last_modification IS NULL) THEN p_description ELSE description END,
        description_last_modification = CASE WHEN (description_last_modification IS NULL) THEN NULL ELSE  description_last_modification END,
        -- virtualization_system
        virtualization_system = CASE WHEN (virtualization_system_last_modification IS NULL) THEN p_virtualization_system ELSE virtualization_system END,
        virtualization_system_last_modification = CASE WHEN (virtualization_system_last_modification IS NULL) THEN NULL ELSE  virtualization_system_last_modification END,
        -- serial_number
        serial_number = CASE WHEN (serial_number_last_modification IS NULL) THEN p_serial_number ELSE serial_number END,
        serial_number_last_modification = CASE WHEN (serial_number_last_modification IS NULL) THEN NULL ELSE  serial_number_last_modification END,
        -- machine_type
        machine_type = CASE WHEN (machine_type_last_modification IS NULL) THEN p_machine_type ELSE machine_type END,
        machine_type_last_modification = CASE WHEN (machine_type_last_modification IS NULL) THEN NULL ELSE  machine_type_last_modification END,
        -- perimeter_id
        perimeter_id = CASE WHEN (perimeter_id_last_modification IS NULL) THEN p_perimeter_id ELSE perimeter_id END,
        perimeter_id_last_modification = CASE WHEN (perimeter_id_last_modification IS NULL) THEN NULL ELSE  perimeter_id_last_modification END,
        -- location_id
        location_id = CASE WHEN (location_id_last_modification IS NULL) THEN p_location_id ELSE location_id END,
        location_id_last_modification = CASE WHEN (location_id_last_modification IS NULL) THEN NULL ELSE  location_id_last_modification END,
        -- operating_system_id
        operating_system_id = CASE WHEN (operating_system_id_last_modification IS NULL) THEN p_operating_system_id ELSE operating_system_id END,
        operating_system_id_last_modification = CASE WHEN (operating_system_id_last_modification IS NULL) THEN NULL ELSE  operating_system_id_last_modification END,
        -- omnis_version
        omnis_version = CASE WHEN (omnis_version_last_modification IS NULL) THEN p_omnis_version ELSE omnis_version END,
        omnis_version_last_modification = CASE WHEN (omnis_version_last_modification IS NULL) THEN NULL ELSE  omnis_version_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE Machine SET
        uuid = COALESCE(p_uuid , uuid),
        authorized = COALESCE(p_authorized , authorized),
        hostname = COALESCE(p_hostname, hostname),
        label = COALESCE(p_label, label),
        description = COALESCE(p_description, description),
        virtualization_system = COALESCE(p_virtualization_system, virtualization_system),
        serial_number = COALESCE(p_serial_number, serial_number),
        machine_type = COALESCE(p_machine_type, machine_type),
        perimeter_id = COALESCE(p_perimeter_id, perimeter_id),
        location_id = COALESCE(p_location_id, location_id),
        operating_system_id = COALESCE(p_operating_system_id, operating_system_id),
        omnis_version = COALESCE(p_omnis_version, omnis_version)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_machine
DROP PROCEDURE IF EXISTS delete_machine//
CREATE PROCEDURE delete_machine(IN p_id INT)
BEGIN
    DELETE FROM Machine WHERE id=p_id;
END //


-- get_installed_softwares
DROP PROCEDURE IF EXISTS get_installed_softwares//
CREATE PROCEDURE get_installed_softwares(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, software_id,machine_id FROM InstalledSoftware WHERE automatic=p_automatic;
END //


-- get_installed_software_by_id
DROP PROCEDURE IF EXISTS get_installed_software_by_id//
CREATE PROCEDURE get_installed_software_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, software_id,machine_id FROM InstalledSoftware WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_installed_softwares
DROP PROCEDURE IF EXISTS get_timestamp_installed_softwares//
CREATE PROCEDURE get_timestamp_installed_softwares(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  software_id_last_modification,machine_id_last_modification
    FROM InstalledSoftware
    WHERE automatic = p_automatic
    AND (software_id_last_modification IS NOT NULL OR machine_id_last_modification IS NOT NULL);
END //


-- get_timestamp_installed_software_by_id
DROP PROCEDURE IF EXISTS get_timestamp_installed_software_by_id//
CREATE PROCEDURE get_timestamp_installed_software_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, software_id_last_modification,machine_id_last_modification
    FROM InstalledSoftware
    WHERE automatic = p_automatic AND id = p_id
    AND (software_id_last_modification IS NOT NULL OR machine_id_last_modification IS NOT NULL);
END //


-- insert_installed_software
DROP PROCEDURE IF EXISTS insert_installed_software//
CREATE PROCEDURE insert_installed_software(IN p_software_id INT,IN p_machine_id INT, IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO InstalledSoftware(automatic, software_id,machine_id,  software_id_last_modification,machine_id_last_modification)
    VALUES(p_automatic, p_software_id,p_machine_id,
    CASE WHEN (p_software_id IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_machine_id IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO InstalledSoftware(id, automatic, software_id,machine_id)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_software_id,p_machine_id);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_installed_software_t
DROP TRIGGER IF EXISTS update_installed_software_t//
CREATE TRIGGER update_installed_software_t BEFORE UPDATE ON InstalledSoftware
FOR EACH ROW BEGIN

IF NEW.software_id != OLD.software_id THEN SET NEW.software_id_last_modification = NOW(); END IF;

IF NEW.machine_id != OLD.machine_id THEN SET NEW.machine_id_last_modification = NOW(); END IF;

END//


-- update_installed_software
DROP PROCEDURE IF EXISTS update_installed_software//
CREATE PROCEDURE update_installed_software(IN p_id INT, IN p_software_id INT,IN p_machine_id INT, IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE InstalledSoftware SET
        software_id = COALESCE(p_software_id, software_id),
        machine_id = COALESCE(p_machine_id, machine_id)
        WHERE id = p_id AND automatic = true;

        UPDATE InstalledSoftware SET
        -- software_id
        software_id = CASE WHEN (software_id_last_modification IS NULL) THEN p_software_id ELSE software_id END,
        software_id_last_modification = CASE WHEN (software_id_last_modification IS NULL) THEN NULL ELSE  software_id_last_modification END,
        -- machine_id
        machine_id = CASE WHEN (machine_id_last_modification IS NULL) THEN p_machine_id ELSE machine_id END,
        machine_id_last_modification = CASE WHEN (machine_id_last_modification IS NULL) THEN NULL ELSE  machine_id_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE InstalledSoftware SET
        software_id = COALESCE(p_software_id, software_id),
        machine_id = COALESCE(p_machine_id, machine_id)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_installed_software
DROP PROCEDURE IF EXISTS delete_installed_software//
CREATE PROCEDURE delete_installed_software(IN p_id INT)
BEGIN
    DELETE FROM InstalledSoftware WHERE id=p_id;
END //


-- get_tagged_machines
DROP PROCEDURE IF EXISTS get_tagged_machines//
CREATE PROCEDURE get_tagged_machines(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, tag_id,machine_id FROM TaggedMachine WHERE automatic=p_automatic;
END //


-- get_tagged_machine_by_id
DROP PROCEDURE IF EXISTS get_tagged_machine_by_id//
CREATE PROCEDURE get_tagged_machine_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, tag_id,machine_id FROM TaggedMachine WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_tagged_machines
DROP PROCEDURE IF EXISTS get_timestamp_tagged_machines//
CREATE PROCEDURE get_timestamp_tagged_machines(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  tag_id_last_modification,machine_id_last_modification
    FROM TaggedMachine
    WHERE automatic = p_automatic
    AND (tag_id_last_modification IS NOT NULL OR machine_id_last_modification IS NOT NULL);
END //


-- get_timestamp_tagged_machine_by_id
DROP PROCEDURE IF EXISTS get_timestamp_tagged_machine_by_id//
CREATE PROCEDURE get_timestamp_tagged_machine_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, tag_id_last_modification,machine_id_last_modification
    FROM TaggedMachine
    WHERE automatic = p_automatic AND id = p_id
    AND (tag_id_last_modification IS NOT NULL OR machine_id_last_modification IS NOT NULL);
END //


-- insert_tagged_machine
DROP PROCEDURE IF EXISTS insert_tagged_machine//
CREATE PROCEDURE insert_tagged_machine(IN p_tag_id INT,IN p_machine_id INT, IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO TaggedMachine(automatic, tag_id,machine_id,  tag_id_last_modification,machine_id_last_modification)
    VALUES(p_automatic, p_tag_id,p_machine_id,
    CASE WHEN (p_tag_id IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_machine_id IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO TaggedMachine(id, automatic, tag_id,machine_id)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_tag_id,p_machine_id);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_tagged_machine_t
DROP TRIGGER IF EXISTS update_tagged_machine_t//
CREATE TRIGGER update_tagged_machine_t BEFORE UPDATE ON TaggedMachine
FOR EACH ROW BEGIN

IF NEW.tag_id != OLD.tag_id THEN SET NEW.tag_id_last_modification = NOW(); END IF;

IF NEW.machine_id != OLD.machine_id THEN SET NEW.machine_id_last_modification = NOW(); END IF;

END//


-- update_tagged_machine
DROP PROCEDURE IF EXISTS update_tagged_machine//
CREATE PROCEDURE update_tagged_machine(IN p_id INT, IN p_tag_id INT,IN p_machine_id INT, IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE TaggedMachine SET
        tag_id = COALESCE(p_tag_id, tag_id),
        machine_id = COALESCE(p_machine_id, machine_id)
        WHERE id = p_id AND automatic = true;

        UPDATE TaggedMachine SET
        -- tag_id
        tag_id = CASE WHEN (tag_id_last_modification IS NULL) THEN p_tag_id ELSE tag_id END,
        tag_id_last_modification = CASE WHEN (tag_id_last_modification IS NULL) THEN NULL ELSE  tag_id_last_modification END,
        -- machine_id
        machine_id = CASE WHEN (machine_id_last_modification IS NULL) THEN p_machine_id ELSE machine_id END,
        machine_id_last_modification = CASE WHEN (machine_id_last_modification IS NULL) THEN NULL ELSE  machine_id_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE TaggedMachine SET
        tag_id = COALESCE(p_tag_id, tag_id),
        machine_id = COALESCE(p_machine_id, machine_id)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_tagged_machine
DROP PROCEDURE IF EXISTS delete_tagged_machine//
CREATE PROCEDURE delete_tagged_machine(IN p_id INT)
BEGIN
    DELETE FROM TaggedMachine WHERE id=p_id;
END //


-- get_networks
DROP PROCEDURE IF EXISTS get_networks//
CREATE PROCEDURE get_networks(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,INET_NTOA(ipv4),ipv4_mask,is_dmz,has_wifi,perimeter_id FROM Network WHERE automatic=p_automatic;
END //


-- get_network_by_id
DROP PROCEDURE IF EXISTS get_network_by_id//
CREATE PROCEDURE get_network_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,INET_NTOA(ipv4),ipv4_mask,is_dmz,has_wifi,perimeter_id FROM Network WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_networks
DROP PROCEDURE IF EXISTS get_timestamp_networks//
CREATE PROCEDURE get_timestamp_networks(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  name_last_modification,ipv4_last_modification,ipv4_mask_last_modification,is_dmz_last_modification,has_wifi_last_modification,perimeter_id_last_modification
    FROM Network
    WHERE automatic = p_automatic
    AND (name_last_modification IS NOT NULL OR ipv4_last_modification IS NOT NULL OR ipv4_mask_last_modification IS NOT NULL OR is_dmz_last_modification IS NOT NULL OR has_wifi_last_modification IS NOT NULL OR perimeter_id_last_modification IS NOT NULL);
END //


-- get_timestamp_network_by_id
DROP PROCEDURE IF EXISTS get_timestamp_network_by_id//
CREATE PROCEDURE get_timestamp_network_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name_last_modification,ipv4_last_modification,ipv4_mask_last_modification,is_dmz_last_modification,has_wifi_last_modification,perimeter_id_last_modification
    FROM Network
    WHERE automatic = p_automatic AND id = p_id
    AND (name_last_modification IS NOT NULL OR ipv4_last_modification IS NOT NULL OR ipv4_mask_last_modification IS NOT NULL OR is_dmz_last_modification IS NOT NULL OR has_wifi_last_modification IS NOT NULL OR perimeter_id_last_modification IS NOT NULL);
END //


-- insert_network
DROP PROCEDURE IF EXISTS insert_network//
CREATE PROCEDURE insert_network(IN p_name VARCHAR(255),IN p_ipv4 VARCHAR(20),IN p_ipv4_mask INT,IN p_is_dmz INT,IN p_has_wifi INT,IN p_perimeter_id INT, IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO Network(automatic, name,ipv4,ipv4_mask,is_dmz,has_wifi,perimeter_id,  name_last_modification,ipv4_last_modification,ipv4_mask_last_modification,is_dmz_last_modification,has_wifi_last_modification,perimeter_id_last_modification)
    VALUES(p_automatic, p_name,INET_ATON(p_ipv4),p_ipv4_mask,p_is_dmz,p_has_wifi,p_perimeter_id,
    CASE WHEN (p_name IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (INET_ATON(p_ipv4) IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_ipv4_mask IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_is_dmz IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_has_wifi IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_perimeter_id IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO Network(id, automatic, name,ipv4,ipv4_mask,is_dmz,has_wifi,perimeter_id)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_name,INET_ATON(p_ipv4),p_ipv4_mask,p_is_dmz,p_has_wifi,p_perimeter_id);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_network_t
DROP TRIGGER IF EXISTS update_network_t//
CREATE TRIGGER update_network_t BEFORE UPDATE ON Network
FOR EACH ROW BEGIN

IF NEW.name != OLD.name THEN SET NEW.name_last_modification = NOW(); END IF;

IF NEW.ipv4 != OLD.ipv4 THEN SET NEW.ipv4_last_modification = NOW(); END IF;

IF NEW.ipv4_mask != OLD.ipv4_mask THEN SET NEW.ipv4_mask_last_modification = NOW(); END IF;

IF NEW.is_dmz != OLD.is_dmz THEN SET NEW.is_dmz_last_modification = NOW(); END IF;

IF NEW.has_wifi != OLD.has_wifi THEN SET NEW.has_wifi_last_modification = NOW(); END IF;

IF NEW.perimeter_id != OLD.perimeter_id THEN SET NEW.perimeter_id_last_modification = NOW(); END IF;

END//


-- update_network
DROP PROCEDURE IF EXISTS update_network//
CREATE PROCEDURE update_network(IN p_id INT, IN p_name VARCHAR(255),IN p_ipv4 VARCHAR(20),IN p_ipv4_mask INT,IN p_is_dmz INT,IN p_has_wifi INT,IN p_perimeter_id INT, IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE Network SET
        name = COALESCE(p_name, name),
        ipv4 = COALESCE(INET_ATON(p_ipv4), ipv4),
        ipv4_mask = COALESCE(p_ipv4_mask, ipv4_mask),
        is_dmz = COALESCE(p_is_dmz, is_dmz),
        has_wifi = COALESCE(p_has_wifi, has_wifi),
        perimeter_id = COALESCE(p_perimeter_id, perimeter_id)
        WHERE id = p_id AND automatic = true;

        UPDATE Network SET
        -- name
        name = CASE WHEN (name_last_modification IS NULL) THEN p_name ELSE name END,
        name_last_modification = CASE WHEN (name_last_modification IS NULL) THEN NULL ELSE  name_last_modification END,
        -- ipv4
        ipv4 = CASE WHEN (ipv4_last_modification IS NULL) THEN INET_ATON(p_ipv4) ELSE ipv4 END,
        ipv4_last_modification = CASE WHEN (ipv4_last_modification IS NULL) THEN NULL ELSE  ipv4_last_modification END,
        -- ipv4_mask
        ipv4_mask = CASE WHEN (ipv4_mask_last_modification IS NULL) THEN p_ipv4_mask ELSE ipv4_mask END,
        ipv4_mask_last_modification = CASE WHEN (ipv4_mask_last_modification IS NULL) THEN NULL ELSE  ipv4_mask_last_modification END,
        -- is_dmz
        is_dmz = CASE WHEN (is_dmz_last_modification IS NULL) THEN p_is_dmz ELSE is_dmz END,
        is_dmz_last_modification = CASE WHEN (is_dmz_last_modification IS NULL) THEN NULL ELSE  is_dmz_last_modification END,
        -- has_wifi
        has_wifi = CASE WHEN (has_wifi_last_modification IS NULL) THEN p_has_wifi ELSE has_wifi END,
        has_wifi_last_modification = CASE WHEN (has_wifi_last_modification IS NULL) THEN NULL ELSE  has_wifi_last_modification END,
        -- perimeter_id
        perimeter_id = CASE WHEN (perimeter_id_last_modification IS NULL) THEN p_perimeter_id ELSE perimeter_id END,
        perimeter_id_last_modification = CASE WHEN (perimeter_id_last_modification IS NULL) THEN NULL ELSE  perimeter_id_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE Network SET
        name = COALESCE(p_name, name),
        ipv4 = COALESCE(INET_ATON(p_ipv4), ipv4),
        ipv4_mask = COALESCE(p_ipv4_mask, ipv4_mask),
        is_dmz = COALESCE(p_is_dmz, is_dmz),
        has_wifi = COALESCE(p_has_wifi, has_wifi),
        perimeter_id = COALESCE(p_perimeter_id, perimeter_id)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_network
DROP PROCEDURE IF EXISTS delete_network//
CREATE PROCEDURE delete_network(IN p_id INT)
BEGIN
    DELETE FROM Network WHERE id=p_id;
END //

-- get_networks_by_ip
DROP PROCEDURE IF EXISTS get_networks_by_ip;
CREATE PROCEDURE get_networks_by_ip(IN p_ipv4 VARCHAR(255), IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,INET_NTOA(ipv4),ipv4_mask,is_dmz,has_wifi,perimeter_id FROM Network WHERE ipv4=INET_ATON(p_ipv4) AND automatic=p_automatic;
END //



-- get_interfaces
DROP PROCEDURE IF EXISTS get_interfaces//
CREATE PROCEDURE get_interfaces(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,INET_NTOA(ipv4),ipv4_mask,mac,interface_type,machine_id,network_id FROM Interface WHERE automatic=p_automatic;
END //


-- get_interface_by_id
DROP PROCEDURE IF EXISTS get_interface_by_id//
CREATE PROCEDURE get_interface_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,INET_NTOA(ipv4),ipv4_mask,mac,interface_type,machine_id,network_id FROM Interface WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_interfaces
DROP PROCEDURE IF EXISTS get_timestamp_interfaces//
CREATE PROCEDURE get_timestamp_interfaces(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  name_last_modification,ipv4_last_modification,ipv4_mask_last_modification,mac_last_modification,interface_type_last_modification,machine_id_last_modification,network_id_last_modification
    FROM Interface
    WHERE automatic = p_automatic
    AND (name_last_modification IS NOT NULL OR ipv4_last_modification IS NOT NULL OR ipv4_mask_last_modification IS NOT NULL OR mac_last_modification IS NOT NULL OR interface_type_last_modification IS NOT NULL OR machine_id_last_modification IS NOT NULL OR network_id_last_modification IS NOT NULL);
END //


-- get_timestamp_interface_by_id
DROP PROCEDURE IF EXISTS get_timestamp_interface_by_id//
CREATE PROCEDURE get_timestamp_interface_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name_last_modification,ipv4_last_modification,ipv4_mask_last_modification,mac_last_modification,interface_type_last_modification,machine_id_last_modification,network_id_last_modification
    FROM Interface
    WHERE automatic = p_automatic AND id = p_id
    AND (name_last_modification IS NOT NULL OR ipv4_last_modification IS NOT NULL OR ipv4_mask_last_modification IS NOT NULL OR mac_last_modification IS NOT NULL OR interface_type_last_modification IS NOT NULL OR machine_id_last_modification IS NOT NULL OR network_id_last_modification IS NOT NULL);
END //


-- insert_interface
DROP PROCEDURE IF EXISTS insert_interface//
CREATE PROCEDURE insert_interface(IN p_name VARCHAR(255),IN p_ipv4 VARCHAR(20),IN p_ipv4_mask INT,IN p_mac VARCHAR(255),IN p_interface_type VARCHAR(255),IN p_machine_id INT,IN p_network_id INT, IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO Interface(automatic, name,ipv4,ipv4_mask,mac,interface_type,machine_id,network_id,  name_last_modification,ipv4_last_modification,ipv4_mask_last_modification,mac_last_modification,interface_type_last_modification,machine_id_last_modification,network_id_last_modification)
    VALUES(p_automatic, p_name,INET_ATON(p_ipv4),p_ipv4_mask,p_mac,p_interface_type,p_machine_id,p_network_id,
    CASE WHEN (p_name IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (INET_ATON(p_ipv4) IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_ipv4_mask IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_mac IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_interface_type IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_machine_id IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_network_id IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO Interface(id, automatic, name,ipv4,ipv4_mask,mac,interface_type,machine_id,network_id)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, p_name,INET_ATON(p_ipv4),p_ipv4_mask,p_mac,p_interface_type,p_machine_id,p_network_id);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_interface_t
DROP TRIGGER IF EXISTS update_interface_t//
CREATE TRIGGER update_interface_t BEFORE UPDATE ON Interface
FOR EACH ROW BEGIN

IF NEW.name != OLD.name THEN SET NEW.name_last_modification = NOW(); END IF;

IF NEW.ipv4 != OLD.ipv4 THEN SET NEW.ipv4_last_modification = NOW(); END IF;

IF NEW.ipv4_mask != OLD.ipv4_mask THEN SET NEW.ipv4_mask_last_modification = NOW(); END IF;

IF NEW.mac != OLD.mac THEN SET NEW.mac_last_modification = NOW(); END IF;

IF NEW.interface_type != OLD.interface_type THEN SET NEW.interface_type_last_modification = NOW(); END IF;

IF NEW.machine_id != OLD.machine_id THEN SET NEW.machine_id_last_modification = NOW(); END IF;

IF NEW.network_id != OLD.network_id THEN SET NEW.network_id_last_modification = NOW(); END IF;

END//


-- update_interface
DROP PROCEDURE IF EXISTS update_interface//
CREATE PROCEDURE update_interface(IN p_id INT, IN p_name VARCHAR(255),IN p_ipv4 VARCHAR(20),IN p_ipv4_mask INT,IN p_mac VARCHAR(255),IN p_interface_type VARCHAR(255),IN p_machine_id INT,IN p_network_id INT, IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE Interface SET
        name = COALESCE(p_name, name),
        ipv4 = COALESCE(INET_ATON(p_ipv4), ipv4),
        ipv4_mask = COALESCE(p_ipv4_mask, ipv4_mask),
        mac = COALESCE(p_mac, mac),
        interface_type = COALESCE(p_interface_type, interface_type),
        machine_id = COALESCE(p_machine_id, machine_id),
        network_id = COALESCE(p_network_id, network_id)
        WHERE id = p_id AND automatic = true;

        UPDATE Interface SET
        -- name
        name = CASE WHEN (name_last_modification IS NULL) THEN p_name ELSE name END,
        name_last_modification = CASE WHEN (name_last_modification IS NULL) THEN NULL ELSE  name_last_modification END,
        -- ipv4
        ipv4 = CASE WHEN (ipv4_last_modification IS NULL) THEN INET_ATON(p_ipv4) ELSE ipv4 END,
        ipv4_last_modification = CASE WHEN (ipv4_last_modification IS NULL) THEN NULL ELSE  ipv4_last_modification END,
        -- ipv4_mask
        ipv4_mask = CASE WHEN (ipv4_mask_last_modification IS NULL) THEN p_ipv4_mask ELSE ipv4_mask END,
        ipv4_mask_last_modification = CASE WHEN (ipv4_mask_last_modification IS NULL) THEN NULL ELSE  ipv4_mask_last_modification END,
        -- mac
        mac = CASE WHEN (mac_last_modification IS NULL) THEN p_mac ELSE mac END,
        mac_last_modification = CASE WHEN (mac_last_modification IS NULL) THEN NULL ELSE  mac_last_modification END,
        -- interface_type
        interface_type = CASE WHEN (interface_type_last_modification IS NULL) THEN p_interface_type ELSE interface_type END,
        interface_type_last_modification = CASE WHEN (interface_type_last_modification IS NULL) THEN NULL ELSE  interface_type_last_modification END,
        -- machine_id
        machine_id = CASE WHEN (machine_id_last_modification IS NULL) THEN p_machine_id ELSE machine_id END,
        machine_id_last_modification = CASE WHEN (machine_id_last_modification IS NULL) THEN NULL ELSE  machine_id_last_modification END,
        -- network_id
        network_id = CASE WHEN (network_id_last_modification IS NULL) THEN p_network_id ELSE network_id END,
        network_id_last_modification = CASE WHEN (network_id_last_modification IS NULL) THEN NULL ELSE  network_id_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE Interface SET
        name = COALESCE(p_name, name),
        ipv4 = COALESCE(INET_ATON(p_ipv4), ipv4),
        ipv4_mask = COALESCE(p_ipv4_mask, ipv4_mask),
        mac = COALESCE(p_mac, mac),
        interface_type = COALESCE(p_interface_type, interface_type),
        machine_id = COALESCE(p_machine_id, machine_id),
        network_id = COALESCE(p_network_id, network_id)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_interface
DROP PROCEDURE IF EXISTS delete_interface//
CREATE PROCEDURE delete_interface(IN p_id INT)
BEGIN
    DELETE FROM Interface WHERE id=p_id;
END //

-- get_interface_by_mac
DROP PROCEDURE IF EXISTS get_interface_by_mac;
CREATE PROCEDURE get_interface_by_mac(IN p_mac VARCHAR(255), IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,INET_NTOA(ipv4),ipv4_mask,mac,interface_type,machine_id,network_id FROM Interface WHERE mac=p_mac AND automatic=p_automatic;
END //


-- get_interfaces_by_machine_id
DROP PROCEDURE IF EXISTS get_interfaces_by_machine_id;
CREATE PROCEDURE get_interfaces_by_machine_id(IN p_machine_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, name,INET_NTOA(ipv4),ipv4_mask,mac,interface_type,machine_id,network_id FROM Interface WHERE machine_id=p_machine_id AND automatic=p_automatic;
END //



-- get_gateways
DROP PROCEDURE IF EXISTS get_gateways//
CREATE PROCEDURE get_gateways(IN p_automatic BOOLEAN)
BEGIN
    SELECT id, INET_NTOA(ipv4),mask,interface_id FROM Gateway WHERE automatic=p_automatic;
END //


-- get_gateway_by_id
DROP PROCEDURE IF EXISTS get_gateway_by_id//
CREATE PROCEDURE get_gateway_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, INET_NTOA(ipv4),mask,interface_id FROM Gateway WHERE id=p_id AND automatic=p_automatic;
END //


-- get_timestamp_gateways
DROP PROCEDURE IF EXISTS get_timestamp_gateways//
CREATE PROCEDURE get_timestamp_gateways(IN p_automatic BOOLEAN)
BEGIN
    SELECT id,  ipv4_last_modification,mask_last_modification,interface_id_last_modification
    FROM Gateway
    WHERE automatic = p_automatic
    AND (ipv4_last_modification IS NOT NULL OR mask_last_modification IS NOT NULL OR interface_id_last_modification IS NOT NULL);
END //


-- get_timestamp_gateway_by_id
DROP PROCEDURE IF EXISTS get_timestamp_gateway_by_id//
CREATE PROCEDURE get_timestamp_gateway_by_id(IN p_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, ipv4_last_modification,mask_last_modification,interface_id_last_modification
    FROM Gateway
    WHERE automatic = p_automatic AND id = p_id
    AND (ipv4_last_modification IS NOT NULL OR mask_last_modification IS NOT NULL OR interface_id_last_modification IS NOT NULL);
END //


-- insert_gateway
DROP PROCEDURE IF EXISTS insert_gateway//
CREATE PROCEDURE insert_gateway(IN p_ipv4 VARCHAR(20),IN p_mask INT,IN p_interface_id INT, IN p_automatic BOOLEAN)
BEGIN
    INSERT INTO Gateway(automatic, ipv4,mask,interface_id,  ipv4_last_modification,mask_last_modification,interface_id_last_modification)
    VALUES(p_automatic, INET_ATON(p_ipv4),p_mask,p_interface_id,
    CASE WHEN (INET_ATON(p_ipv4) IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_mask IS NULL) THEN NULL ELSE NOW() END,
    CASE WHEN (p_interface_id IS NULL) THEN NULL ELSE NOW() END);

    INSERT INTO Gateway(id, automatic, ipv4,mask,interface_id)
    VALUES(LAST_INSERT_ID(), NOT p_automatic, INET_ATON(p_ipv4),p_mask,p_interface_id);

    SELECT LAST_INSERT_ID() AS id;
END //


-- update_gateway_t
DROP TRIGGER IF EXISTS update_gateway_t//
CREATE TRIGGER update_gateway_t BEFORE UPDATE ON Gateway
FOR EACH ROW BEGIN

IF NEW.ipv4 != OLD.ipv4 THEN SET NEW.ipv4_last_modification = NOW(); END IF;

IF NEW.mask != OLD.mask THEN SET NEW.mask_last_modification = NOW(); END IF;

IF NEW.interface_id != OLD.interface_id THEN SET NEW.interface_id_last_modification = NOW(); END IF;

END//


-- update_gateway
DROP PROCEDURE IF EXISTS update_gateway//
CREATE PROCEDURE update_gateway(IN p_id INT, IN p_ipv4 VARCHAR(20),IN p_mask INT,IN p_interface_id INT, IN p_automatic BOOLEAN)
BEGIN
    if p_automatic THEN  -- auto
        UPDATE Gateway SET
        ipv4 = COALESCE(INET_ATON(p_ipv4), ipv4),
        mask = COALESCE(p_mask, mask),
        interface_id = COALESCE(p_interface_id, interface_id)
        WHERE id = p_id AND automatic = true;

        UPDATE Gateway SET
        -- ipv4
        ipv4 = CASE WHEN (ipv4_last_modification IS NULL) THEN INET_ATON(p_ipv4) ELSE ipv4 END,
        ipv4_last_modification = CASE WHEN (ipv4_last_modification IS NULL) THEN NULL ELSE  ipv4_last_modification END,
        -- mask
        mask = CASE WHEN (mask_last_modification IS NULL) THEN p_mask ELSE mask END,
        mask_last_modification = CASE WHEN (mask_last_modification IS NULL) THEN NULL ELSE  mask_last_modification END,
        -- interface_id
        interface_id = CASE WHEN (interface_id_last_modification IS NULL) THEN p_interface_id ELSE interface_id END,
        interface_id_last_modification = CASE WHEN (interface_id_last_modification IS NULL) THEN NULL ELSE  interface_id_last_modification END
        WHERE id = p_id AND automatic = false;
    ELSE -- manual
        UPDATE Gateway SET
        ipv4 = COALESCE(INET_ATON(p_ipv4), ipv4),
        mask = COALESCE(p_mask, mask),
        interface_id = COALESCE(p_interface_id, interface_id)
        WHERE id = p_id AND automatic = false;
    END IF;
END //


-- delete_gateway
DROP PROCEDURE IF EXISTS delete_gateway//
CREATE PROCEDURE delete_gateway(IN p_id INT)
BEGIN
    DELETE FROM Gateway WHERE id=p_id;
END //

-- get_gateways_by_interface_id
DROP PROCEDURE IF EXISTS get_gateways_by_interface_id;
CREATE PROCEDURE get_gateways_by_interface_id(IN p_interface_id INT, IN p_automatic BOOLEAN)
BEGIN
    SELECT id, INET_NTOA(ipv4),mask,interface_id FROM Gateway WHERE interface_id=p_interface_id AND automatic=p_automatic;
END //

DELIMITER ; //
