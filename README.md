# How To Run
1. Clone repository
2. Pastikan docker berjalan
3. Buka folder project
4. Buka Terminal di folder project, jalankan perintah command ini 
   ```console
        docker run -d -v ${PWD}/envoy.yaml:/etc/envoy/envoy.yaml:ro -p 8080:8080 -p 9901:9901 envoyproxy/envoy:v1.22.0

        docker build -t go-grpc-server:1.0 -f server/Dockerfile server

        docker run -d  -p 9090:9090 --env-file server/.env go-grpc-server:1.0 
   ```
5. Jalankan docker ps, atau cek di docker desktop apakah ada 2 container baru yang berhasil dijalankan
6. Pindah ke folder client
7. Jalankan "yarn" 
8. Jalankan "yarn start"
9.  Happy Coding!
