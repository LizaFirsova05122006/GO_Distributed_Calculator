from flask import Flask, render_template, request
import configure as cfg
import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart
import sqlite3
import logging

app = Flask(__name__)

# Настройка логирования
logging.basicConfig(level=logging.DEBUG)

# Конфигурация сообщения
msg = MIMEMultipart()
msg["From"] = cfg.LOGIN

@app.route("/forgot", methods=['POST', 'GET'])
def question():
    if request.method == 'GET':
        return render_template("forgot.html")
    elif request.method == 'POST':
        try:
            param = {}
            param['email'] = request.form['email']

            connection = sqlite3.connect('users.db')
            cursor = connection.cursor()
            cursor.execute('SELECT token, password FROM users WHERE email == ?', (param['email'],))
            users = cursor.fetchall()

            if not users:
                return "Пользователь с данным email не найден", 404

            tokens = users[0][0]
            passw = users[0][1]

            msg["Subject"] = "Ваши данные для входа"
            msg["To"] = param['email']
            msg_body = (f"Ваши данные для входа:\n" 
                        f"JWT Token: {tokens}\n"
                        f"Пароль: {passw}")
            msg.attach(MIMEText(msg_body, "plain"))

            # Настройка SMTP-сервера
            server = smtplib.SMTP_SSL("smtp.yandex.ru", 465)
            server.set_debuglevel(1)  # Включение отладочного режима

            # Логин и отправка
            server.login(cfg.LOGIN, cfg.PASSWORD)
            server.sendmail(cfg.LOGIN, param['email'], msg.as_string())
            server.quit()

            connection.close()

            return render_template('succes.html', title="Успешно")
        except smtplib.SMTPException as e:
            logging.error(f"Ошибка при отправке email: {e}")
            return f"Ошибка при отправке email: {e}", 500
        except sqlite3.Error as e:
            logging.error(f"Ошибка базы данных: {e}")
            return f"Ошибка базы данных: {e}", 500

if __name__ == "__main__":
    app.run()
