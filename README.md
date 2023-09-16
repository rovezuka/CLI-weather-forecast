# Weather CLI application

Это CLI (Command Line Interface) приложение, разработанное на Go, позволяет вам получать актуальную погодную информацию для любого города в мире и просматривать прогноз погоды на весь день.

## Описание

Погодное CLI приложение предоставляет простой и удобный способ получить информацию о погоде в интересующем вас городе. Вы можете узнать текущую температуру, погодные условия и вероятность дождя для выбранного местоположения. Кроме того, приложение также предоставляет подробный прогноз погоды на весь день, включая температуру на разные часы и условия погоды.

## Требования

Для использования этого приложения вам потребуется:

- Установленный Go (версия 1.16 и выше).
- API ключ от провайдера погоды. Вы можете получить его, зарегистрировавшись на сайте провайдера погоды.

## Получение API ключа

1. Перейдите на [сайт провайдера погоды](https://www.weatherapi.com/signup.aspx).
2. Зарегистрируйтесь или войдите в свой аккаунт.
3. Создайте новый API ключ.
4. Скопируйте полученный ключ.

## Установка и запуск

Подробные инструкции по установке и запуску вашего CLI приложения.

1. Склонируйте репозиторий:
```
git clone https://github.com/rovezuka/CLI-application.git

```

2. Перейдите в директорию проекта:
```
cd CLI-application

```

3. Установите зависимости:
```
go get github.com/fatih/color

```

4. В HTTP-запросе замените API в параметре "key" на личный

5. Скомпилируйте проект:
```
go build
```

6. Запустите CLI приложение в командной строке:
```
CLI-application Moscow 

```

## Использование

Примеры использования:

- Получить прогноз погоды в Москве:
```
CLI-application Moscow 
```
- Получить прогноз погоды в Нью-Йорке:
```
CLI-application New-York 
```
- Получить прогноз погоды в Лондоне:
```
CLI-application London 
```

## Вклад

Если вы хотите внести свой вклад в развитие этого проекта, вы можете сделать следующее:

Откройте задачу (issue), если у вас есть предложения или баги для улучшения.
Создайте pull request с вашими изменениями, если вы хотите добавить новые функции или улучшения.

## Лицензия

Этот проект лицензируется в соответствии с лицензией MIT. Вы можете прочитать полный текст лицензии в файле LICENSE.

Создано с ❤️ с использованием Go.

Этот README.md предоставляет базовую информацию о проекте и как начать его использовать. Вы можете дополнить его более подробными сведениями о структуре проекта, настройке, документации по API, примерами использования и др.

