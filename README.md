<h1 align='center'>Проект "Распределенный вычислитель арифметических выражений"</h1>
<br></br>
<nav>
  <h2>Содержание</h2>
  <ol>
    <li><a href="#section1">Подготовка к запуску проекта</a></li>
    <ul>
      <li><a href="#section23">Установка дополнительных библиотек</a></li>
      <ul>
        <li><a href="#section26">github.com/mattn/go-sqlite3</a></li>
        <li><a href="#section27">github.com/jinzhu/gorm</a></li>
        <li><a href="#section28">github.com/gin-gonic/gin</a></li>
        <li><a href="#section29">github.com/dgrijalva/jwt-go</a></li>
        <li><a href="#section30">github.com/gorilla/mux</a></li>
        <li><a href="#section31">github.com/Knetic/govaluate</a></li>
      </ul>
      <li><a href="#section24">Установка компилятора GCC</a></li>
      <li><a href="#section25">Добавление переменных в launch.json</a></li>
    </ul>
    <li><a href="#section7">Использованные технологии</a></li>
    <ul>
      <li><a href="#section2">Яндекс Почта</a></li>
      <li><a href="#section6">Программа на языке программирования Python</a></li>
      <li><a href="#section8">Базы данных</a></li>
      <ul>
        <li><a href="#section9">users.db</a></li>
        <li><a href="#section10">rezults.db</a></li>
      </ul>
    </ul>
    <li><a href="#section3">Техническое описание проекта</a></li>
    <li><a href="#section4">Инструкция по настройке и запуску проекта</a></li>
    <li><a href="#section5">Реализация проекта через сайт</a></li>
    <ul>
      <li><a href="#section15">Registration Form</a></li>
      <li><a href="#section16">Login Form</a></li>
      <li><a href="#section17">Forgot the JWT</a></li>
      <li><a href="#section18">User Profile</a></li>
      <li><a href="#section19">Data entry Form</a></li>
      <li><a href="#section20">Results</a></li>
    </ul>
    <li><a href="#section11">Реализация проекта через командную строку</a></li>
    <ul>
      <li><a href="#section12">Регистрация пользователя в системе</a></li>
      <ul>
        <li><a href="#section13">Успешная регистрация</a></li>
        <li><a href="#section14">Ошибки при регистрации</a></li>
        <ul>
          <li><a href="#section21">Логин уже существует</a></li>
          <li><a href="#section22">Введены пустые данные</a></li>
        </ul>
     </ul>
    <li><a href="#section36">Вход пользователя в систему</a></li>
      <ul>
        <li><a href="#section32">Успешный вход</a></li>
        <li><a href="#section33">Ошибки при входе</a></li>
        <ul>
          <li><a href="#section34">Неверные логин или пароль</a></li>
          <li><a href="#section35">Введены пустые данные</a></li>
       </ul>
    </ul>
    </ul>
  </ol>
</nav>
<br></br>
<h1 id="section1">Подготовка к запуску проекта</h1>
<h1 id="section23">Установка дополнительных библиотек</h1>
<h1 id="section26">github.com/mattn/go-sqlite3</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/DZRhCQMS/2024-07-10-08-28-11.png' border='0' alt='2024-07-10-08-28-11'/></a>
<h1 id="section27">github.com/jinzhu/gorm</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/qq5JWHxt/2024-07-11-16-22-03.png' border='0' alt='2024-07-11-16-22-03'/></a>
<h1 id="section28">github.com/gin-gonic/gin</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/W4v2Bng7/2024-07-11-16-22-34.png' border='0' alt='2024-07-11-16-22-34'/></a>
<h1 id="section29">github.com/golang-jwt/jwt/v5</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/7LqGnzC1/2024-07-10-10-51-54.png' border='0' alt='2024-07-10-10-51-54'/></a>
<h1 id="section30">github.com/gorilla/mux</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/wjhzQmZR/2024-07-11-16-39-17.png' border='0' alt='2024-07-11-16-39-17'/></a>
<h1 id="section31">github.com/Knetic/govaluate</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/5N8wbdL2/2024-07-11-11-25-39.png' border='0' alt='2024-07-11-11-25-39'/></a>
<h1 id="section24">Установка компилятора GCC</h1>
<h4>Необходимо установить компилятор GCC с сайта <a href="https://jmeubank.github.io/tdm-gcc/">Tdm-gcc</a>:</h4>
<ul>
 <li>Нажать кнопку Create и соглашаться со всем</li>
 <li>Запомнить папку, куда установилось</li>
</ul>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/T18wyXGN/2024-07-10-08-34-29.png' border='0' alt='2024-07-10-08-34-29'/></a>
<h1 id="section25">Добавление переменных в launch.json</h1>
<h4>3) В VS Code открыть Run and Debug и создать файл launch.json:</h4>
<ul>
  <li>Добавить переменную CGO_ENABLED=1</li>
  <li>В перменную Path добавить путь до gcc.exe</li>
</ul>  
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/vH0skdkR/2024-07-10-08-41-32.png' border='0' alt='2024-07-10-08-41-32'/></a>
<br></br>
<h1 id="section7">Использованные технологии</h1>
<h1 id="section2">Яндекс Почта</h1>
<h4>Для того, чтобы отправлять пользователям сообщения был создан отдельнный почтовый ящик. В приложении Яндекс Почты был создан пароль приложений. Чтобы точно все работало в настройках необходимо поставить галочки:</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/TPkcN47Q/2024-07-10-17-31-28.png' border='0' alt='2024-07-10-17-31-28'/></a>
<br></br>
<h1 id="section6">Программа на языке программирования Python</h1>
<h4>Для того, чтобы отправлять письма пользователям, была написана программа на языке программирования Python, процесс:</h4>
<ol>
  <li>Пользователь вводит адрес электронной почты</li>
  <li>Программа забирает email и делает запрос в базу данных на проверку существует ли такой email в базе данных</li>
  <li>Забирает по email password и JWT Token</li>
  <li>Формирует письмо</li>
  <li>Регистрируется на сайте при помощь логина и пароля записанных в файле configure.py</li>
  <li>Отправляет письмо на email</li>
</ol>
<br></br>
<h1 id="section8">Базы данных</h1>
<h4>В проекте используются две базы данных</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/y8Pp6Scx/2024-07-11-11-54-39.png' border='0' alt='2024-07-11-11-54-39'/></a>
<h1 id="section9">users.db</h1>
<h4>База данных, которая используется на этапе регистрации и входа в систему.</h4>
<h1 id="section10">rezults.db</h1>
<h4>База данных, которая используется на этапе обработки выражений. В нее вносятся выражения и результаты. Также из нее экспортируются данные для отображения на странице результатов.</h4>
<br></br>
<h1 id="section3">Техническое описание проекта</h1>
<ol>
  <li>Templates - папка с шаблонами:<ul>
    <li>forgot.html - шаблон страницы, где пользователь вводит свой email. Программа проверяет наличие введенного email в базе данных, затем отправляет письмо</li>
    <li>succes.html - шаблон страницы, после нажатии кнопки отправить</li>
  </ul></li>
  <li>Configure.py - файл, где хранится логин и пароль от почтового ящика, с которого отправляются письма</li>
  <li>Forgot.py - файл, в котором написана программа страницы, где пользователь вводит email</li>
</ol>
<br></br>
<h1 id="section4">Инструкция по настройке и запуску проекта</h1>
<h4>Перед тем как запускать проект, скачайте все файлы в один проект. Для настройки и запуска проекта необходимо создать виртуальное окружение.</h4>
<br></br>
<h1 id="section5">Реализация проекта через сайт</h1>
<h4>Перед началом проекта рекомендуется запустить register.go через Visual Studio Code и forgot.py через PyCharm.</h4>
<h1 id="section15">Registration Form</h1>
<h4>Перейдите на сайт http://localhost:8080/register если Вы впервые заходите на сайт:</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/rpTkn9yY/2024-07-10-22-40-52.png' border='0' alt='2024-07-10-22-40-52'/></a>
<h4>Введите адрес электронной почты, пароль, повторите пароль. После регистрации будет выведен JWT Token, а также дальнейшая инструкция</h4>
<h1 id="section16">Login Form</h1>
<h4>Перейдите на сайт http://localhost:8080/login если у Вас уже есть JWT Token:</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/Z5Sb2mVc/2024-07-10-22-47-16.png' border='0' alt='2024-07-10-22-47-16'/></a>
<h4>Введите JWT Token для входа в систему. Также можно создать аккаунт, если у Вас нет аккаунта; восстановить JWT Token</h4>
<h1 id="section17">Forgot the JWT</h1>
<h4>Перейдя на сайт http://127.0.0.1:5000/forgot необходимо ввести адрес электронной почты. Затем программа проверит наличия адреса электронной почты в базе данных. После нажатия кнопки Отправить программа формирует письмо, отправляет его пользователю</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/hP4MrQqb/2024-07-10-22-53-54.png' border='0' alt='2024-07-10-22-53-54'/></a>
<h1 id="section18">User Profile</h1>
<h4>Перейдя на сайт http://localhost:8080/user/{email}?email={email}, вместо {email} программа подставляет реальнный адрес электронной почты. Профиль пользователя:</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/Y2y75Cc6/2024-07-11-11-29-21.png' border='0' alt='2024-07-11-11-29-21'/></a>
<h1 id="section19">Data entry Form</h1>
<h4>Перейдя на сайт(или нажав на кнопку "Ввести выражение") http:localhost:8080/calculate/{email}?email={email}, куда вместо {email} программа подставляет адрес электронной почты человека, можно написать выражение для вычисления. После нажатия кноки "Отправить" программа добавить выражение в базу данных.</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/vTmvD1vL/2024-07-11-10-24-08.png' border='0' alt='2024-07-11-10-24-08'/></a>
<h1 id="section20">Results</h1>
<h4>Перейдя на сайт(или нажав на кнопку "Проверить результаты") http:localhost:8080/rezults/{email}?email={email}, вместо {email} программа подставляет реальнный адрес электронной почты. На странице можно проверить результаты выражений введенных пользователем:</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/PqB3W908/2024-07-11-11-38-04.png' border='0' alt='2024-07-11-11-38-04'/></a>
<br></br>
<h1 id="section11">Реализация проекта через командную строку</h1>
<h1 id="section12">Регистрация пользователя в системе</h1>
<h1 id="section13">Успешная регистрация</h1>
<h4>Для запуска запустите терминал. Для примера использовался Windows PowerShell</h4>
<ol>
  <li><h4>Заголовки. Функция, задающая контент-тип для тела запроса как application/json</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/7ZsmD9rh/2024-07-11-15-36-36.png' border='0' alt='2024-07-11-15-36-36'/></a>
  <li><h4>Тело запроса. Здесь Вы задаете JSON содержимое для тела запроса, представляющее собой логин и пароль пользователя</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/pTdLqYkt/2024-07-14-17-05-10.png' border='0' alt='2024-07-14-17-05-10'/></a>
  <li><h4>Преобразование в JSON. При помощи ConvertTo-Json тело запроса преобразуется в JSON строку.</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/PJMQTKRn/2024-07-11-15-46-23.png' border='0' alt='2024-07-11-15-46-23'/></a>
  <li><h4>Выполнение запроса. Invoke-WebRequest выполняет POST-запрос с указанными заголовками и телом.</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/pL3zNgLn/2024-07-11-15-49-00.png' border='0' alt='2024-07-11-15-49-00'/></a>
  <li><h4>Вывод ответа. Вывод содержимого ответа на экране</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/NMDK0VBD/2024-07-11-15-50-52.png' border='0' alt='2024-07-11-15-50-52'/></a>
</ol>
<h4>В результате нам вышло сообщение об успешной регистрации:</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/JnnPM0RW/2024-07-14-19-15-46.png' border='0' alt='2024-07-14-19-15-46'/></a>
<h1 id="section14">Ошибки при регистрации</h1>
<h1 id="section21">Логин уже существует</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/FRqNGt7j/2024-07-14-17-08-31.png' border='0' alt='2024-07-14-17-08-31'/></a>
<h1 id="section22">Введены пустые данные</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/Y2GW7MfB/2024-07-14-17-09-35.png' border='0' alt='2024-07-14-17-09-35'/></a>
<h1 id="section36">Вход пользователя в систему</h1>
<h1 id="section32">Успешный вход</h1>
<ol>
  <li><h4>Заголовки. Функция, задающая контент-тип для тела запроса как application/json</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/7ZsmD9rh/2024-07-11-15-36-36.png' border='0' alt='2024-07-11-15-36-36'/></a>
  <li><h4>Тело запроса. Здесь Вы задаете JSON содержимое для тела запроса, представляющее собой логин и пароль пользователя</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/2jTLnGdT/2024-07-14-19-19-11.png' border='0' alt='2024-07-14-19-19-11'/></a>
  <li><h4>Преобразование в JSON. При помощи ConvertTo-Json тело запроса преобразуется в JSON строку.</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/PJMQTKRn/2024-07-11-15-46-23.png' border='0' alt='2024-07-11-15-46-23'/></a>
  <li><h4>Выполнение запроса. Invoke-WebRequest выполняет POST-запрос с указанными заголовками и телом.</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/05H28B0g/2024-07-14-19-20-04.png' border='0' alt='2024-07-14-19-20-04'/></a>
  <li><h4>Вывод ответа. Вывод содержимого ответа на экране</h4></li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/NMDK0VBD/2024-07-11-15-50-52.png' border='0' alt='2024-07-11-15-50-52'/></a>
</ol>
<h4>В результате нам вышло сообщение об успешной регистрации:</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/nLkpNRhP/2024-07-14-19-21-01.png' border='0' alt='2024-07-14-19-21-01'/></a>
<h1 id="section33">Ошибки при входе</h1>
<h1 id="section34">Неверные логин или пароль</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/y8KN5STW/2024-07-14-19-24-18.png' border='0' alt='2024-07-14-19-24-18'/></a>
<h1 id="section35">Введены пустые данные</h1>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/vZPdz1qJ/2024-07-14-19-25-29.png' border='0' alt='2024-07-14-19-25-29'/></a>
