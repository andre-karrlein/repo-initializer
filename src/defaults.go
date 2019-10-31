package main

func getDefaultDirectories() []string {
	return []string{
		"config",
		"docker",
		"docker/nginx",
		"docker/php-fpm",
		"public",
		"src",
		"test",
	}
}

func getDefaultFiles() []string {
	return []string{
		"config/integration.ini",
		"config/staging.ini",
		"config/production.ini",
		"docker/nginx/Dockerfile",
		"docker/php-fpm/Dockerfile",
		"public/index.php",
		"src/Factory.php",
		"test/FactoryTest.php",
		"docker-compose.yml",
		"docker-compose-dev.yml",
		"docker-compose-prd.yml",
		"README.md",
	}
}
