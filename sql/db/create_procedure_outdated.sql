USE OMNIS;

DELIMITER // ;

-- retrieve outdated perimeters
CREATE OR REPLACE PROCEDURE get_outdated_perimeters(IN p_outdated_day INT)
BEGIN
    SELECT * FROM Perimeter WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        description_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated locations
CREATE OR REPLACE PROCEDURE get_outdated_locations(IN p_outdated_day INT)
BEGIN
    SELECT * FROM Location WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        description_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated OperatingSystems
CREATE OR REPLACE PROCEDURE get_outdated_operating_systems(IN p_outdated_day INT)
BEGIN
    SELECT * FROM OperatingSystem WHERE automatic=false AND (
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
    SELECT * FROM Tag WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        color_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated softwares
CREATE OR REPLACE PROCEDURE get_outdated_softwares(IN p_outdated_day INT)
BEGIN
    SELECT * FROM Software WHERE automatic=false AND (
        name_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        version_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        is_intern_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated machines
CREATE OR REPLACE PROCEDURE get_outdated_machines(IN p_outdated_day INT)
BEGIN
    SELECT * FROM Machine WHERE automatic=false AND (
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
    SELECT * FROM InstalledSoftware WHERE automatic=false AND (
        software_id_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        machine_id_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated TaggedMachines
CREATE OR REPLACE PROCEDURE get_outdated_tagged_machines(IN p_outdated_day INT)
BEGIN
    SELECT * FROM TaggedMachine WHERE automatic=false AND (
        tag_id_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        machine_id_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //

-- retrieve outdated networks
CREATE OR REPLACE PROCEDURE get_outdated_networks(IN p_outdated_day INT)
BEGIN
    SELECT * FROM Network WHERE automatic=false AND (
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
    SELECT * FROM Interface WHERE automatic=false AND (
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
    SELECT * FROM Gateway WHERE automatic=false AND (
        ipv4_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        mask_last_modification < NOW() - INTERVAL p_outdated_day DAY OR
        interface_id_last_modification < NOW() - INTERVAL p_outdated_day DAY
    );
END //


DELIMITER ; //
