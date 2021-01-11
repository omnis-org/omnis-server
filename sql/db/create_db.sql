DROP DATABASE IF EXISTS OMNIS;
CREATE DATABASE OMNIS;
GRANT ALL PRIVILEGES ON OMNIS.* TO 'omnis'@'localhost' IDENTIFIED BY 'MyBeautifulPassword8273!';


DROP DATABASE IF EXISTS OMNIS_ADMIN;
CREATE DATABASE OMNIS_ADMIN;
GRANT ALL PRIVILEGES ON OMNIS_ADMIN.* TO 'omnis'@'localhost' IDENTIFIED BY 'MyBeautifulPassword8273!';

-- Create omnis database to store data about machines and network

USE OMNIS;

--------------------------------------------------------------------------

-- Explain automatic :

-- The automatic field in the tables shows where the information comes from.
-- An attribute _last_modification is also added in the tables to know when the data was updated.

-- When inserting, two tables are created but only the timestamp of the table with the boolean of the request is updated.

-- For example: :

-- insert_o(name = "abc" , desc = NULL, automatic = true)
-- o_a = "abc" , NOW() , NULL, NULL
-- o_m = "abc" , NULL, NULL, NULL

-- insert_o(name = "abc", desc = NULL , automatic = false)
-- o_a = "abc" , NULL, NULL, NULL
-- o_m = "abc" , NOW(), NULL, NULL


-- When we update a table manually, we only change the manual table and these timestamps.

-- When we update a table automatically, we modify the automatic table and these timestamps.
-- But we also modify the attributes not modified in the manual table but not the timestamps.


-- For example: :

-- update_o(name = "def", automatic = true)
-- o_a = "def", NOW()

----

-- (old o_m1 = "abc", NULL)
-- o_m1 = "def", NULL

----

-- (old o_m2 = "abc", OLDNOW)
-- o_m2 = "abc", OLDNOW

-------------

-- update_o(name = "def", automatic = false)
-- o_m = "def", NOW()

--------------------------------------------------------------------------

-- Perimeter
CREATE TABLE Perimeter (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    name_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    description_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT perimeter_name_uq UNIQUE (name, automatic),
    CONSTRAINT perimeter_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;

-- Location
CREATE TABLE Location (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    name_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    description_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT location_name_uq UNIQUE (name, automatic),
    CONSTRAINT location_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;

CREATE TABLE OperatingSystem (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    platform VARCHAR(100),
    platform_family VARCHAR(100),
    platform_version VARCHAR(100),
    kernel_version VARCHAR(100),
    name_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    platform_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    platform_family_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    platform_version_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    kernel_version_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT operating_system_uq UNIQUE (name, platform, platform_family, platform_version, kernel_version, automatic),
    CONSTRAINT operating_system_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;

-- CREATE TABLE MachineType (
--     id INT NOT NULL AUTO_INCREMENT,
--     name VARCHAR(255) NOT NULL,
--     automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
--     CONSTRAINT machine_type_name_uq UNIQUE (name, automatic),
--     CONSTRAINT machine_type_pk PRIMARY KEY (id, automatic)
-- ) WITH SYSTEM VERSIONING;

CREATE TABLE Tag (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(10) NOT NULL,
    name_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    color_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT tag_name_uq UNIQUE (name, automatic),
    CONSTRAINT tag_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;

CREATE TABLE Software (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    version VARCHAR(255),
    is_intern BOOLEAN default 0,
    name_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    version_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    is_intern_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT software_name_version_uq UNIQUE (name, version, automatic),
    CONSTRAINT software_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;


CREATE TABLE Machine (
    id INT NOT NULL AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL,
    authorized BOOLEAN DEFAULT NULL,
    hostname VARCHAR(255) NOT NULL,
    label VARCHAR(255) NOT NULL,
    description TEXT, -- 2^16 - 1
    virtualization_system VARCHAR(255),
    serial_number VARCHAR(255),
    machine_type ENUM('client','server','router'),
    perimeter_id INT NOT NULL,
    location_id INT NOT NULL,
    operating_system_id INT,
    omnis_version VARCHAR(255),
    uuid_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    authorized_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    hostname_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    label_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    description_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    virtualization_system_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    serial_number_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    machine_type_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    perimeter_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    location_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    operating_system_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    omnis_version_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT machine_uuid_uq UNIQUE (uuid, automatic),
    CONSTRAINT machine_perimeter_fk FOREIGN KEY (perimeter_id) REFERENCES Perimeter(id),
    CONSTRAINT machine_location_fk FOREIGN KEY (location_id) REFERENCES Location(id),
    CONSTRAINT machine_operating_system_fk FOREIGN KEY (operating_system_id) REFERENCES OperatingSystem(id),
    CONSTRAINT machine_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;


CREATE TABLE InstalledSoftware (
    id INT NOT NULL AUTO_INCREMENT,
    software_id INT NOT NULL,
    machine_id INT NOT NULL,
    software_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    machine_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT installed_software_software_machine_uq UNIQUE (software_id, machine_id, automatic),
    CONSTRAINT installed_software_software_fk FOREIGN KEY (software_id) REFERENCES Software(id),
    CONSTRAINT installed_software_machine_fk FOREIGN KEY (machine_id) REFERENCES Machine(id),
    CONSTRAINT installed_software_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;

CREATE TABLE TaggedMachine (
    id INT NOT NULL AUTO_INCREMENT,
    tag_id INT NOT NULL,
    machine_id INT NOT NULL,
    tag_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    machine_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT tagged_machine_tag_machine_uq UNIQUE (tag_id, machine_id, automatic),
    CONSTRAINT tagged_machine_tag_fk FOREIGN KEY (tag_id) REFERENCES Tag(id),
    CONSTRAINT tagged_machine_machine_fk FOREIGN KEY (machine_id) REFERENCES Machine(id),
    CONSTRAINT tagged_machine_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;

-- ------------------------------------------------------------------------
-- INET_ATON ; INET_NTOA ; https://stackoverflow.com/questions/2542011/most-efficient-way-to-store-ip-address-in-mysql

CREATE TABLE Network (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    ipv4 INT UNSIGNED NOT NULL,
    ipv4_mask INT NOT NULL,
    is_dmz BOOLEAN default 0,
    has_wifi BOOLEAN default 0,
    perimeter_id INT NOT NULL,
    name_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    ipv4_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    ipv4_mask_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    is_dmz_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    has_wifi_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    perimeter_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT gateway_network_ipv4_ipv4_mask_perimeter_id_uq UNIQUE (ipv4, ipv4_mask, perimeter_id, automatic),
    CONSTRAINT network_perimeter_fk FOREIGN KEY (perimeter_id) REFERENCES Perimeter(id),
    CONSTRAINT network_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;

-- ------------------------------------------------------------------------
-- INET_ATON ; INET_NTOA


-- CREATE TABLE InterfaceType(
--     id INT NOT NULL AUTO_INCREMENT,
--     name VARCHAR(255) NOT NULL,
--     automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
--     CONSTRAINT interface_type_name_uq UNIQUE (name, automatic),
--     CONSTRAINT interface_type_pk PRIMARY KEY (id, automatic)
-- ) WITH SYSTEM VERSIONING;


CREATE TABLE Interface (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    ipv4 INT UNSIGNED NOT NULL,
    ipv4_mask INT NOT NULL,
    mac VARCHAR(255),
    interface_type ENUM("eth", "wan", "lo"),
    machine_id INT NOT NULL,
    network_id INT NOT NULL,
    name_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    ipv4_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    ipv4_mask_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    mac_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    interface_type_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    machine_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    network_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT mac_uq UNIQUE (mac, automatic),
    CONSTRAINT interface_machine_fk FOREIGN KEY (machine_id) REFERENCES Machine(id) ON DELETE CASCADE,
    CONSTRAINT interface_network_fk FOREIGN KEY (network_id) REFERENCES Network(id),
    CONSTRAINT interface_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;

CREATE TABLE Gateway(
    id INT NOT NULL AUTO_INCREMENT,
    ipv4 INT UNSIGNED NOT NULL,
    mask INT NOT NULL,
    interface_id INT NOT NULL,
    ipv4_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    mask_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    interface_id_last_modification TIMESTAMP NULL DEFAULT NULL INVISIBLE WITHOUT SYSTEM VERSIONING,
    automatic BOOLEAN NOT NULL WITHOUT SYSTEM VERSIONING,
    CONSTRAINT gateway_ipv4_mask_interface_uq UNIQUE (ipv4, mask, interface_id, automatic),
    CONSTRAINT gateway_interface_fk FOREIGN KEY (interface_id) REFERENCES Interface(id) ON DELETE CASCADE,
    CONSTRAINT gateway_pk PRIMARY KEY (id, automatic)
) WITH SYSTEM VERSIONING;

-- ------------------------------------------------------------------------




USE OMNIS_ADMIN;

-- PERMISSION
-- 0000 = NO PERMSSION = 0;
-- 0001 = SELECT = 1;
-- 0010 = INSERT = 2;
-- 0100 = UPDATE = 4;
-- 1000 = DELETE = 8;
-- 1111 = ALL PERMSSION = 15;

CREATE TABLE Role(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    omnis_permissions INT NOT NULL,
    roles_permissions INT NOT NULL,
    users_permissions INT NOT NULL,
    pending_machines_permissions INT NOT NULL,
    CONSTRAINT role_pk PRIMARY KEY (id)
);

INSERT INTO Role VALUES(1,"Administrator",15,15,15,15);

CREATE TABLE User(
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(64) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    role_id INT NOT NULL,
    CONSTRAINT username_uq UNIQUE (username),
    CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES Role(id),
    CONSTRAINT user_pk PRIMARY KEY (id)
);
