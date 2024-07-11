<h1 align='center'>Проект "Распределенный вычислитель арифметических выражений"</h1>
<br></br>
<nav>
  <h2>Содержание</h2>
  <ol>
    <li><a href="#section1">Подготовка к запуску проекта</a></li>
    <li><a href="#section7">Использованные технологии</a></li>
    <ul>
      <li><a href="#section2">Яндекс Почта</a></li>
      <li><a href="#section6">Программа на языке программирования Python</a></li>
    </ul>
    <li><a href="#section3">Техническое описание проекта</a></li>
    <li><a href="#section4">Инструкция по настройке и запуску проекта</a></li>
    <li><a href="#section5">Реализация проекта через сайт</a></li>
  </ol>
</nav>
<br></br>
<h1 id="section1">Подготовка к запуску проекта</h1>
<h4>1) Перед началом необходимо установить "github.com/mattn/go-sqlite3". Для этого откроем Терминал:</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/DZRhCQMS/2024-07-10-08-28-11.png' border='0' alt='2024-07-10-08-28-11'/></a>
<h4>2) Также необходимо установить компилятор GCC с сайта <a href="https://jmeubank.github.io/tdm-gcc/">Tdm-gcc</a>:</h4>
<ul>
 <li>Нажать кнопку Create и соглашаться со всем</li>
 <li>Запомнить папку, куда установилось</li>
</ul>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/T18wyXGN/2024-07-10-08-34-29.png' border='0' alt='2024-07-10-08-34-29'/></a>
<h4>3) В VS Code открыть Run and Debug и создать файл launch.json:</h4>
<ul>
  <li>Добавить переменную CGO_ENABLED=1</li>
  <li>В перменную Path добавить путь до gcc.exe</li>
</ul>  
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/vH0skdkR/2024-07-10-08-41-32.png' border='0' alt='2024-07-10-08-41-32'/></a>
<h4>4) Необходимо в командной строке написать: "go get -u github.com/golang-jwt/jwt/v5" для установки библиотеки:</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/7LqGnzC1/2024-07-10-10-51-54.png' border='0' alt='2024-07-10-10-51-54'/></a>
<h4>5)Установить библиотеку при помощи команды: "go get github.com/Knetic/govaluate":</h4>
<a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/5N8wbdL2/2024-07-11-11-25-39.png' border='0' alt='2024-07-11-11-25-39'/></a>
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
<ol>
  <li>Перейдите на сайт http://localhost:8080/register если Вы впервые заходите на сайт:</li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/rpTkn9yY/2024-07-10-22-40-52.png' border='0' alt='2024-07-10-22-40-52'/></a>
  <h4>Введите адрес электронной почты, пароль, повторите пароль. После регистрации будет выведен JWT Token, а также дальнейшая инструкция</h4>
  <li>Перейдите на сайт http://localhost:8080/login если у Вас уже есть JWT Token:</li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/Z5Sb2mVc/2024-07-10-22-47-16.png' border='0' alt='2024-07-10-22-47-16'/></a>
  <h4>Введите JWT Token для входа в систему. Также можно создать аккаунт, если у Вас нет аккаунта; восстановить JWT Token</h4>
  <li>Перейдя на сайт http://127.0.0.1:5000/forgot необходимо ввести адрес электронной почты. Затем программа проверит наличия адреса электронной почты в базе данных. После нажатия кнопки Отправить программа формирует письмо, отправляет его пользователю</li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/hP4MrQqb/2024-07-10-22-53-54.png' border='0' alt='2024-07-10-22-53-54'/></a>
  <li>Перейдя на сайт http://localhost:8080/user/{email}?email={email}, вместо {email} программа подставляет реальнный адрес электронной почты. Профиль пользователя:</li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/Y2y75Cc6/2024-07-11-11-29-21.png' border='0' alt='2024-07-11-11-29-21'/></a>
  <li>Перейдя на сайт(или нажав на кнопку "Ввести выражение") http:localhost:8080/calculate/{email}?email={email}, куда вместо {email} программа подставляет адрес электронной почты человека, можно написать выражение для вычисления. После нажатия кноки "Отправить" программа добавить выражение в базу данных.</li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/vTmvD1vL/2024-07-11-10-24-08.png' border='0' alt='2024-07-11-10-24-08'/></a>
  <li>Перейдя на сайт(или нажав на кнопку "Проверить результаты") http:localhost:8080/rezults/{email}?email={email}, вместо {email} программа подставляет реальнный адрес электронной почты. На странице можно проверить результаты выражений введенных пользователем:</li>
  <a href='https://postimages.org/' target='_blank'><img src='https://i.postimg.cc/PqB3W908/2024-07-11-11-38-04.png' border='0' alt='2024-07-11-11-38-04'/></a>
</ol>
