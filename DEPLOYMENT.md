# 🚀 Руководство по развертыванию Space API

## Содержание
1. [Локальное развертывание](#локальное-развертывание)
2. [Docker развертывание](#docker-развертывание)
3. [Облачное развертывание](#облачное-развертывание)
4. [Настройка CI/CD](#настройка-cicd)
5. [Мониторинг и логирование](#мониторинг-и-логирование)

---

## Локальное развертывание

### Вариант 1: Готовый бинарник

1. **Скачайте релиз** для вашей платформы:
   ```bash
   # Пример для Linux
   wget https://github.com/your-repo/releases/download/v1.0.0/space-api-linux-amd64
   chmod +x space-api-linux-amd64
   ```

2. **Запустите:**
   ```bash
   ./space-api-linux-amd64
   ```

3. **Настройка порта (опционально):**
   ```bash
   export PORT=3000
   ./space-api-linux-amd64
   ```

### Вариант 2: Сборка из исходников

1. **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/your-repo/go-fake-api.git
   cd go-fake-api
   ```

2. **Установите зависимости:**
   ```bash
   go mod download
   ```

3. **Соберите:**
   ```bash
   go build -ldflags="-s -w" -o space-api main.go
   ```

4. **Запустите:**
   ```bash
   ./space-api
   ```

### Запуск как системный сервис (Linux)

1. **Создайте systemd service файл:**
   ```bash
   sudo nano /etc/systemd/system/space-api.service
   ```

2. **Добавьте конфигурацию:**
   ```ini
   [Unit]
   Description=Space API Service
   After=network.target

   [Service]
   Type=simple
   User=www-data
   WorkingDirectory=/opt/space-api
   ExecStart=/opt/space-api/space-api
   Restart=always
   RestartSec=10
   Environment="PORT=8080"

   [Install]
   WantedBy=multi-user.target
   ```

3. **Запустите сервис:**
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable space-api
   sudo systemctl start space-api
   sudo systemctl status space-api
   ```

---

## Docker развертывание

### Вариант 1: Docker Compose (рекомендуется)

1. **Создайте docker-compose.yml:**
   ```yaml
   version: '3.8'
   
   services:
     space-api:
       build: .
       ports:
         - "8080:8080"
       environment:
         - PORT=8080
       restart: unless-stopped
   ```

2. **Запустите:**
   ```bash
   docker-compose up -d
   ```

3. **Проверьте логи:**
   ```bash
   docker-compose logs -f
   ```

### Вариант 2: Docker напрямую

1. **Соберите образ:**
   ```bash
   docker build -t space-api:latest .
   ```

2. **Запустите контейнер:**
   ```bash
   docker run -d \
     --name space-api \
     -p 8080:8080 \
     -e PORT=8080 \
     --restart unless-stopped \
     space-api:latest
   ```

3. **Проверьте статус:**
   ```bash
   docker ps
   docker logs space-api
   ```

### Использование готового образа из Docker Hub

```bash
docker pull your-username/space-api:latest
docker run -d -p 8080:8080 your-username/space-api:latest
```

---

## Облачное развертывание

### AWS EC2

1. **Создайте EC2 инстанс** (Ubuntu 22.04 LTS)

2. **Подключитесь по SSH:**
   ```bash
   ssh -i your-key.pem ubuntu@your-ec2-ip
   ```

3. **Установите Docker:**
   ```bash
   sudo apt update
   sudo apt install -y docker.io docker-compose
   sudo systemctl enable docker
   sudo usermod -aG docker ubuntu
   ```

4. **Клонируйте и запустите:**
   ```bash
   git clone https://github.com/your-repo/go-fake-api.git
   cd go-fake-api
   docker-compose up -d
   ```

5. **Настройте Security Group:**
   - Откройте порт 8080 для входящих соединений

### Google Cloud Platform (GCP)

1. **Создайте VM инстанс:**
   ```bash
   gcloud compute instances create space-api \
     --image-family=ubuntu-2204-lts \
     --image-project=ubuntu-os-cloud \
     --machine-type=e2-micro \
     --zone=us-central1-a
   ```

2. **Подключитесь:**
   ```bash
   gcloud compute ssh space-api
   ```

3. **Установите и запустите** (см. AWS EC2 шаги 3-4)

4. **Настройте firewall:**
   ```bash
   gcloud compute firewall-rules create allow-space-api \
     --allow tcp:8080 \
     --source-ranges 0.0.0.0/0
   ```

### DigitalOcean

1. **Создайте Droplet** (Ubuntu 22.04)

2. **Используйте Docker One-Click App** или установите вручную

3. **Деплой через Docker** (см. раздел Docker)

### Heroku

1. **Создайте Procfile:**
   ```
   web: ./space-api
   ```

2. **Деплой:**
   ```bash
   heroku create your-space-api
   git push heroku main
   ```

### Railway.app

1. **Подключите GitHub репозиторий**

2. **Railway автоматически определит Dockerfile**

3. **Настройте переменные окружения:**
   - `PORT` (Railway предоставит автоматически)

---

## Настройка CI/CD

### GitHub Actions (уже настроено)

Workflow автоматически:
- Собирает релизы при создании тега
- Создает GitHub Release
- Загружает бинарники для всех платформ
- Собирает и публикует Docker образ

**Использование:**
```bash
git tag v1.0.0
git push origin v1.0.0
```

### GitLab CI/CD

Создайте `.gitlab-ci.yml`:
```yaml
stages:
  - build
  - release

build:
  stage: build
  image: golang:1.21
  script:
    - go build -ldflags="-s -w" -o space-api main.go
  artifacts:
    paths:
      - space-api

release:
  stage: release
  only:
    - tags
  script:
    - echo "Creating release"
```

---

## Настройка обратного прокси

### Nginx

1. **Установите Nginx:**
   ```bash
   sudo apt install nginx
   ```

2. **Создайте конфигурацию:**
   ```bash
   sudo nano /etc/nginx/sites-available/space-api
   ```

3. **Добавьте:**
   ```nginx
   server {
       listen 80;
       server_name your-domain.com;

       location / {
           proxy_pass http://localhost:8080;
           proxy_http_version 1.1;
           proxy_set_header Upgrade $http_upgrade;
           proxy_set_header Connection 'upgrade';
           proxy_set_header Host $host;
           proxy_cache_bypass $http_upgrade;
           proxy_set_header X-Real-IP $remote_addr;
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
           proxy_set_header X-Forwarded-Proto $scheme;
       }
   }
   ```

4. **Активируйте:**
   ```bash
   sudo ln -s /etc/nginx/sites-available/space-api /etc/nginx/sites-enabled/
   sudo nginx -t
   sudo systemctl restart nginx
   ```

### SSL с Let's Encrypt

```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d your-domain.com
sudo systemctl reload nginx
```

---

## Мониторинг и логирование

### Базовый мониторинг

**Проверка здоровья:**
```bash
curl http://localhost:8080/api/gagarin-flight
```

**Мониторинг процесса:**
```bash
ps aux | grep space-api
top -p $(pgrep space-api)
```

### Логирование

**Просмотр логов (systemd):**
```bash
sudo journalctl -u space-api -f
```

**Просмотр логов (Docker):**
```bash
docker logs -f space-api
```

### Prometheus + Grafana (расширенный мониторинг)

Добавьте в будущем:
- Метрики запросов
- Время ответа
- Использование ресурсов

---

## Масштабирование

### Горизонтальное масштабирование

1. **Запустите несколько инстансов:**
   ```bash
   docker run -d -p 8081:8080 space-api:latest
   docker run -d -p 8082:8080 space-api:latest
   docker run -d -p 8083:8080 space-api:latest
   ```

2. **Настройте load balancer** (Nginx, HAProxy)

### Kubernetes

Создайте `deployment.yaml`:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: space-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: space-api
  template:
    metadata:
      labels:
        app: space-api
    spec:
      containers:
      - name: space-api
        image: your-username/space-api:latest
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: space-api
spec:
  selector:
    app: space-api
  ports:
  - port: 80
    targetPort: 8080
  type: LoadBalancer
```

Деплой:
```bash
kubectl apply -f deployment.yaml
```

---

## Безопасность

### Рекомендации

1. **Измените JWT secret:**
   - Не используйте дефолтный secret в продакшене
   - Используйте переменные окружения

2. **Настройте HTTPS:**
   - Используйте SSL сертификаты
   - Редирект с HTTP на HTTPS

3. **Firewall:**
   ```bash
   sudo ufw allow 22/tcp
   sudo ufw allow 80/tcp
   sudo ufw allow 443/tcp
   sudo ufw enable
   ```

4. **Rate limiting:**
   - Настройте в Nginx или используйте middleware

5. **Регулярные обновления:**
   ```bash
   sudo apt update && sudo apt upgrade
   ```

---

## Резервное копирование

### Backup данных (если используется БД)

```bash
# Пример для будущего использования с PostgreSQL
pg_dump space_api > backup.sql
```

### Backup конфигурации

```bash
tar -czf space-api-config-$(date +%Y%m%d).tar.gz \
  /etc/systemd/system/space-api.service \
  /etc/nginx/sites-available/space-api
```

---

## Troubleshooting

### Проблема: Порт занят

```bash
# Найти процесс на порту 8080
sudo lsof -i :8080
# Или
sudo netstat -tulpn | grep 8080

# Убить процесс
sudo kill -9 <PID>
```

### Проблема: Нет доступа

1. Проверьте firewall
2. Проверьте Security Groups (AWS/GCP)
3. Проверьте логи

### Проблема: Высокое использование памяти

```bash
# Перезапустите сервис
sudo systemctl restart space-api
```

---

## Чек-лист деплоя

- [ ] Сервер настроен и обновлен
- [ ] Docker установлен (если используется)
- [ ] Приложение собрано/скачано
- [ ] Порты открыты
- [ ] Nginx настроен (если используется)
- [ ] SSL сертификат установлен
- [ ] Systemd service создан
- [ ] Автозапуск настроен
- [ ] Логирование работает
- [ ] Мониторинг настроен
- [ ] Backup настроен
- [ ] Документация обновлена

---

**Готово!** 🚀 Ваш Space API развернут и готов к работе!
