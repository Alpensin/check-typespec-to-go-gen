# check-typespec-to-go-gen

# Подготовка
Установка компилятора
`npm install -g @typespec/compiler`
# Создание спек
1. Инициализация проекта. Выбран openapi3 generic проект
`tsp init`
2. Скопирован код в main.tsp из https://github.com/microsoft/typespec для примера
3. `tsp install` - создает node_modules
4. Генерация спеки openapi3
`tsp compile main.tsp --emit @typespec/openapi3`

# Генерация кода go.
## Попробовал
1. https://github.com/Azure/autorest.go - он генерит код с привязкой к своему sdk. По этой причине не очень подходит
2. https://github.com/ogen-go/ogen
Запуск с дефолтным конфигом:
```sh
 docker run --rm \
  --volume ".:/workspace" \
  ghcr.io/ogen-go/ogen:latest --target workspace/ogen-gen-code --clean workspace/tsp-output/@typespec/openapi3/openapi.yaml
```
С такой командой ogen генерит много кода для запуска сервера, клиента, валидации клиента и сервера и OpenTelemetry.
Можно настраивать в конфиге
`.ogen.yml`
https://github.com/ogen-go/ogen/blob/main/examples/_config/example_all.yml
даже с выключенными фичами генерит код с привязкой к либам
```
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
```

Понравился https://github.com/oapi-codegen/oapi-codegen
позволяет генерить модели и ручки для разных серверов включая net/http
можно задать конфиг. 
Самый простой вариант. Есть много опций для развития и встраивания свагера.
```
# Название пакета в сгенеренном файле
package: testapapi
# Файл куда сохранятся результаты
output: testapapi.gen.go
generate:
  # генерить модели
  models: true
  # генерить ручки для net/http сервера
  std-http-server: true
```
Поддерживаемые серверы  https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#generating-server-side-boilerplate
```sh
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
oapi-codegen --config=oapiconf.yaml tsp-output/@typespec/openapi3/openapi.yaml
```
Спека по конфигу https://github.com/oapi-codegen/oapi-codegen/blob/main/configuration-schema.json
