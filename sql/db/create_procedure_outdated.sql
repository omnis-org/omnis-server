USE OMNIS;

DELIMITER // ;

-- retrieve outdated perimeters
CREATE OR REPLACE PROCEDURE get_outdated_perimeters(IN p_outdated_day INT)
BEGIN
    SELECT id,name,description,name_last_modification,description_last_modification FROM Perimeter WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        description_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated locations
CREATE OR REPLACE PROCEDURE get_outdated_locations(IN p_outdated_day INT)
BEGIN
    SELECT id,name,description,name_last_modification,description_last_modification FROM Location WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        description_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated OperatingSystems
CREATE OR REPLACE PROCEDURE get_outdated_operating_systems(IN p_outdated_day INT)
BEGIN
    SELECT id,name,platform,platform_family,platform_version,kernel_version,
    name_last_modification,platform_last_modification,platform_family_last_modification,platform_version_last_modification,kernel_version_last_modification
    FROM OperatingSystem WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        platform_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        platform_family_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        platform_version_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        kernel_version_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated tags
CREATE OR REPLACE PROCEDURE get_outdated_tags(IN p_outdated_day INT)
BEGIN
    SELECT id,name,color,name_last_modification,color_last_modification FROM Tag WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        color_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated softwares
CREATE OR REPLACE PROCEDURE get_outdated_softwares(IN p_outdated_day INT)
BEGIN
    SELECT id,name,version,is_intern,name_last_modification,version_last_modification,is_intern_last_modification
    FROM Software WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        version_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        is_intern_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated machines
CREATE OR REPLACE PROCEDURE get_outdated_machines(IN p_outdated_day INT)
BEGIN
    SELECT id,uuid,hostname,label,description,virtualization_system,serial_number,machine_type,perimeter_id,location_id,operating_system_id,omnis_version,
    uuid_last_modification,hostname_last_modification,label_last_modification,description_last_modification,virtualization_system_last_modification,serial_number_last_modification,
    machine_type_last_modification,perimeter_id_last_modification,location_id_last_modification,operating_system_id_last_modification,omnis_version_last_modification
    FROM Machine WHERE automatic=false AND authorized=true AND (
        uuid_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        hostname_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        label_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        description_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        virtualization_system_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        serial_number_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        machine_type_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        perimeter_id_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        location_id_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        operating_system_id_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        omnis_version_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated InstalledSoftwares
CREATE OR REPLACE PROCEDURE get_outdated_installed_softwares(IN p_outdated_day INT)
BEGIN
    SELECT id,software_id,machine_id,software_id_last_modification,machine_id_last_modification
    FROM InstalledSoftware WHERE automatic=false AND (
        software_id_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        machine_id_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated TaggedMachines
CREATE OR REPLACE PROCEDURE get_outdated_tagged_machines(IN p_outdated_day INT)
BEGIN
    SELECT id,tag_id,machine_id,tag_id_last_modification,machine_id_last_modification
    FROM TaggedMachine WHERE automatic=false AND (
        tag_id_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        machine_id_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated networks
CREATE OR REPLACE PROCEDURE get_outdated_networks(IN p_outdated_day INT)
BEGIN
    SELECT id,name,ipv4,ipv4_mask,is_dmz,has_wifi,perimeter_id,
    name_last_modification,ipv4_last_modification,ipv4_mask_last_modification,
    is_dmz_last_modification,has_wifi_last_modification,perimeter_id_last_modification
    FROM Network WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        ipv4_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        ipv4_mask_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        is_dmz_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        has_wifi_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        perimeter_id_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated interfaces
CREATE OR REPLACE PROCEDURE get_outdated_interfaces(IN p_outdated_day INT)
BEGIN
    SELECT id,name,ipv4,ipv4_mask,mac,interface_type,machine_id,network_id
    name_last_modification,ipv4_last_modification,ipv4_mask_last_modification,
    mac_last_modification,interface_type_last_modification,machine_id_last_modification,network_id_last_modification
    FROM Interface WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        ipv4_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        ipv4_mask_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        mac_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        interface_type_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        machine_id_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        network_id_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated gateways
CREATE OR REPLACE PROCEDURE get_outdated_gateways(IN p_outdated_day INT)
BEGIN
    SELECT id,ipv4,mask,interface_id,ipv4_last_modification
    mask_last_modification,interface_id_last_modification
    FROM Gateway WHERE automatic=false AND (
        ipv4_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        mask_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        interface_id_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //


DELIMITER ; //
