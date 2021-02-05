USE OMNIS;
CALL insert_perimeter("SI-BUREAUTIQUE","Système d'information dédié à l'utilisation bureautique",1);
CALL insert_perimeter("SI-BACKUP","Système d'information dédié aux sauvegardes et décorélé des autres SI",1);
CALL insert_location("Paris", "PCA",1);
CALL insert_location("Lyon", "Siège",1);
CALL insert_location("Nomade", "Utilisateur",1);
CALL insert_operating_system("linux", "ubuntu", "debian", "20.04", "5.4.0-50-generic",1);
CALL insert_operating_system("windows", "windows", "windows", "10", "unknown",1);
CALL insert_operating_system("windows", "windows", "server", "2019", "unknown",1);
CALL insert_machine("58c9df97-1bd8-4270-8ace-3177fff1b28e",1,"AD01","AD01","ActiveDirectory domain controller",NULL,"012345","client",1,2,3,"1.0", 1);
CALL insert_machine("2f0096ea-2f02-4882-ad7a-e24ad8c13b5d",1,"FILE-SHARE01","FILE-SHARE01","File sharing server",NULL,"012445","client",1,2,3, "1.0", 1);
CALL insert_machine("ce9de293-2dd2-4e92-a18b-50c0868148f6",1,"AD01-PCA","AD01-PCA","ActiveDirectory domain controller redondant",NULL,"012346","client",1,1,3, "1.0", 1);
CALL insert_machine("558f79e5-3ccb-42f7-964d-16244e550505",1,"FILE-SHARE01-PCA","FILE-SHARE-PCA01","File sharing server","Docker","012347","client",1,1,3, "1.0", 1);
CALL insert_machine("781b0cca-b347-4505-a7d3-af1888cf4fa6",1,"BACKUP01-DOMAIN","BACKUP01-DOMAIN","machine de backup décorélée des autres SI","Podman","29931","client",2,2,1,NULL, 1);
CALL insert_machine("e73ea760-0ab8-48de-a868-6a98bb7f4fae",1,"BACKUP01-DOMAIN-PCA","BACKUP02-DOMAIN-PCA","machine de backup dédiée au site de PCA et décorélée des autres SI","Virtualbox","29431","client",2,1,3,NULL, 1);
CALL insert_machine("5e4f8b41-59c3-407a-aacc-6a7ee8582fcb",NULL,"user01","user01","Domain user computer",NULL,"012315","client",1,3,2, "1.0", 1);
CALL insert_machine("bc4b0aac-7fdb-49c8-918c-12917a031790",NULL,"user02","user02","Domain user computer",NULL,"012325","client",1,3,2, "1.0", 1);
CALL insert_machine("0d82a3eb-8c74-4013-b6bf-df50afd4c7cb",NULL,"user03","user03","Domain user computer",NULL,"012335","client",1,3,2, "1.0", 1);
CALL insert_machine("07012778-9e05-44db-91a0-270ef76caa02",NULL,"user04","user04","Domain user computer",NULL,"012345","client",1,3,2, "1.0", 1);
CALL insert_machine("7f3533c1-2a1c-4d3a-b238-66ac8ed15296",0,"user05","user05","Domain user computer",NULL,"012355","client",1,3,2, "1.0", 1);

CALL update_machine(1,NULL,NULL,NULL, NULL,"Update Description ADO1",NULL,NULL,NULL,NULL,NULL,NULL,NULL,0);
UPDATE Machine SET description_last_modification='2020-01-01 01:01:01' WHERE id=1 AND automatic=false;

CALL update_machine(2,NULL,NULL,NULL, NULL,"Update Description FILE-SHARE",NULL,NULL,NULL,NULL,NULL,NULL,NULL,0);
UPDATE Machine SET description_last_modification='2020-10-10 10:10:10' WHERE id=2 AND automatic=false;


CALL insert_network("LAN-DOMAIN-CONTROL", "192.168.10.0", 24, 0, 0, 1, 1);
CALL insert_network("LAN-DOMAIN-CONTROL-PCA", "172.16.10.0", 24, 0, 0, 1, 1);
CALL insert_network("LAN-DOMAIN-CONTROL-BACKUP", "192.168.11.0", 24, 0, 0, 2, 1);
CALL insert_network("LAN-DOMAIN-CONTROL-PCA-BACKUP", "172.16.11.0", 24, 0, 0, 2, 1);
CALL insert_network("LAN-BUREAUTIQUE", "192.168.20.0", 24, 0, 0, 1, 1);
CALL insert_interface("eth0", "192.168.10.1", 24, "AB:0D:1F", "eth", 1, 1,1);
CALL insert_interface("eth0", "192.168.10.2", 24, "AB:1D:1F", "eth", 2, 1,1);
CALL insert_interface("eth0", "172.16.10.1", 24, "AB:0D:2F", "eth", 3, 2,1);
CALL insert_interface("eth0", "172.16.10.2", 24, "AB:1D:2F", "eth", 4, 2,1);
CALL insert_interface("eth0", "192.168.11.1", 24, "AB:0D:3F", "eth", 5, 3,1);
CALL insert_interface("eth0", "172.16.11.1", 24, "AB:0D:4F", "eth", 6, 4,1);
CALL insert_interface("eth0", "192.168.20.1", 24, "AB:0D:5F", "eth", 7, 5,1);
CALL insert_interface("eth0", "192.168.20.2", 24, "AB:0D:6F", "eth", 8, 5,1);
CALL insert_interface("eth0", "192.168.20.3", 24, "AB:0D:7F", "eth", 9, 5,1);
CALL insert_interface("eth0", "192.168.20.4", 24, "AB:0D:8F", "eth", 10, 5,1);
CALL insert_interface("eth0", "192.168.20.5", 24, "AB:0D:9F", "eth", 11, 5,1);

