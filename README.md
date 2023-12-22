# sdl-go

docker build -t robertcretu/graphapp:latest . -f deploy/Dockerfile
argocd admin initial-password -n argocd
docker push robertcretu/graphapp:latest   
kubectl port-forward svc/graphapp-service -n default 8000:8000
kubectl port-forward svc/argocd-server -n argocd 8080:443
mysql -u root -p -h 127.0.0.1 devops
SELECT * FROM projects go;
SELECT * FROM repositories go;