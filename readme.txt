powershell
backup docker volume 

1.docker run --rm -v postgres_data:/volume -v D:\Benz\Mikelopster\Go-Database:/backup busybox sh -c "tar czf /backup/docker_vol_postgres_data.tar.gz -C /volume ."
2.docker run --rm -v pgadmin_data:/volume -v D:\Benz\Mikelopster\Go-Database:/backup busybox sh -c "tar czf /backup/docker_vol_pgadmin_data.tar.gz -C /volume ."

restore docker volume 
1.docker run --rm -v postgres_data:/volume -v D:/Benz/Mikelopster/Go-Database:/backup busybox sh -c "tar xzf /backup/docker_vol_postgres_data.tar.gz -C /volume"
2.docker run --rm -v pgadmin_data:/volume -v D:/Benz/Mikelopster/Go-Database:/backup busybox sh -c "tar xzf /backup/docker_vol_pgadmin_data.tar.gz -C /volume"

คำสั่ง query หรือ queryRow จะดึงข้อมูลจาก database ออกมาแล้ว จะเป็นการเก็บ pointer ของ row  แต่ยังไม่ทำการโหลดเข้าสู่ memory ต้องใช้ คำสั่ง scan ถึงจะโหลดเข้า memory ทีละตัวๆ