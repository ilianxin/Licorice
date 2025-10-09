# Licorice - Virtual Currency Trading System

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Spring Cloud](https://img.shields.io/badge/Spring%20Cloud-2023.0-green.svg)](https://spring.io/projects/spring-cloud)
[![Java](https://img.shields.io/badge/Java-17-orange.svg)](https://openjdk.java.net/)

## Overview

Licorice is a comprehensive virtual currency trading system built with Spring Cloud microservices architecture. It provides a robust, scalable, and secure platform for cryptocurrency trading operations with a frontend-backend separation design.

## Features

- **Microservices Architecture**: Built on Spring Cloud for high scalability and maintainability
- **Real-time Trading**: Live market data and instant order execution
- **Security**: Advanced security measures for user authentication and transaction protection
- **Multi-currency Support**: Support for various cryptocurrencies
- **User Management**: Comprehensive user account and profile management
- **Trading Engine**: High-performance order matching and execution system
- **API Gateway**: Centralized API management and routing
- **Monitoring**: Real-time system monitoring and alerting

## Technology Stack

### Backend
- **Framework**: Spring Boot 3.x
- **Microservices**: Spring Cloud 2023.0
- **Service Discovery**: Eureka/Nacos
- **API Gateway**: Spring Cloud Gateway
- **Database**: MySQL 8.0, Redis
- **Message Queue**: RabbitMQ/Apache Kafka
- **Security**: Spring Security, JWT
- **Monitoring**: Spring Boot Actuator, Micrometer

### Frontend
- **Framework**: Vue.js 3.x / React 18.x
- **UI Library**: Element Plus / Ant Design
- **State Management**: Pinia / Redux Toolkit
- **Build Tool**: Vite / Webpack
- **Charts**: TradingView / ECharts

### DevOps
- **Containerization**: Docker, Docker Compose
- **Orchestration**: Kubernetes
- **CI/CD**: GitHub Actions / Jenkins
- **Monitoring**: Prometheus, Grafana
- **Logging**: ELK Stack (Elasticsearch, Logstash, Kibana)

## Project Structure

```
licorice/
├── licorice-gateway/          # API Gateway
├── licorice-auth/             # Authentication Service
├── licorice-user/             # User Management Service
├── licorice-trading/          # Trading Engine Service
├── licorice-market/           # Market Data Service
├── licorice-wallet/           # Wallet Service
├── licorice-notification/     # Notification Service
├── licorice-admin/            # Admin Service
├── licorice-common/           # Common Libraries
├── licorice-frontend/         # Frontend Application
├── docker-compose.yml         # Docker Compose Configuration
├── k8s/                       # Kubernetes Manifests
└── docs/                      # Documentation
```

## Quick Start

### Prerequisites

- Java 17+
- Node.js 16+
- Docker & Docker Compose
- MySQL 8.0+
- Redis 6.0+

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/ilianxin/Licorice.git
   cd Licorice
   ```

2. **Start infrastructure services**
   ```bash
   docker-compose up -d mysql redis rabbitmq
   ```

3. **Start backend services**
   ```bash
   # Start service discovery
   cd licorice-eureka && mvn spring-boot:run
   
   # Start API Gateway
   cd licorice-gateway && mvn spring-boot:run
   
   # Start other services
   cd licorice-auth && mvn spring-boot:run
   cd licorice-user && mvn spring-boot:run
   cd licorice-trading && mvn spring-boot:run
   ```

4. **Start frontend**
   ```bash
   cd licorice-frontend
   npm install
   npm run dev
   ```

### Docker Deployment

```bash
# Build and start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Scale services
docker-compose up -d --scale licorice-trading=3
```

## API Documentation

- **Gateway**: http://localhost:8080
- **Swagger UI**: http://localhost:8080/swagger-ui.html
- **API Docs**: http://localhost:8080/v3/api-docs

## Configuration

### Environment Variables

```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_NAME=licorice
DB_USERNAME=root
DB_PASSWORD=password

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT Configuration
JWT_SECRET=your-secret-key
JWT_EXPIRATION=86400000

# External APIs
BINANCE_API_URL=https://api.binance.com
COINBASE_API_URL=https://api.coinbase.com
```

## Security

- **Authentication**: JWT-based authentication
- **Authorization**: Role-based access control (RBAC)
- **Data Encryption**: AES-256 encryption for sensitive data
- **API Security**: Rate limiting and request validation
- **Audit Logging**: Comprehensive audit trail for all operations

## Monitoring

- **Health Checks**: Spring Boot Actuator endpoints
- **Metrics**: Prometheus metrics collection
- **Dashboards**: Grafana visualization
- **Alerting**: AlertManager for critical events
- **Logging**: Centralized logging with ELK stack

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Support

- **Documentation**: [Wiki](https://github.com/ilianxin/Licorice/wiki)
- **Issues**: [GitHub Issues](https://github.com/ilianxin/Licorice/issues)
- **Discussions**: [GitHub Discussions](https://github.com/ilianxin/Licorice/discussions)

## Roadmap

- [ ] Mobile application (React Native)
- [ ] Advanced trading features (futures, options)
- [ ] DeFi integration
- [ ] Multi-language support
- [ ] Advanced analytics and reporting
- [ ] Social trading features

---

**Disclaimer**: This is a demo project for educational purposes. Please ensure compliance with local regulations when dealing with cryptocurrency trading systems.
