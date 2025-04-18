package cogs

import (
	"gopkg.in/telebot.v3"
)

func ScheduleCogs(bot *telebot.Bot) {
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		if c.Text() == "Расписание" {
			return c.Send("Напишите день недели (пн, вт, ср, чт, пт, сб, вс)")
		}

		switch c.Text() {
		case "пн":
			return c.Send(`📅 Понедельник

8:30-10:00
(1-15 неделя) Операционные системы. Амер Исмаил, ауд. 910 (Кремлёвская, 35)

10:10-11:40
(1-15 неделя) Дискретная математика. Пшеничный П.В., ауд. 112 (Кремлёвская, 16А)

13:50-15:20
(1-15 неделя) Организация и правовое обеспечение информационной безопасности. Ситников С.Ю., ауд. 112 (Кремлёвская, 16А)

15:50-17:20
(ч/н, 2-14 неделя) Организация и правовое обеспечение информационной безопасности. Панищев О.Ю., ауд. 1407 (Кремлёвская, 16А)`)
		
		case "вт":
			return c.Send(`📅 Вторник

10:10-11:40
(1-9 неделя) Технологии и методы программирования. Шаймухаметов Р.Р., ауд. 305 (Кремлёвская, 16А)

12:10-13:40
(1-9 неделя) Документоведение. Рубцова Р.Г., ауд. 112 (Кремлёвская, 16А)

15:50-17:20
(7-15 неделя) Деловой иностранный язык для программистов. Ситникова А.С. - в формате ЭОР`)
		
		case "ср":
			return c.Send(`📅 Среда

8:30-10:00
(4-6 неделя) Операционные системы. Амер Исмаил, ауд. 810 (Кремлёвская, 35)
(7-15 неделя) Основы управленческой деятельности. Васильева М.А., ауд. 1117 (Кремлёвская, 16А)

10:10-11:40
(ч/н, 13-15 неделя) Дискретная математика. Пшеничный П.В., ауд. 1113 (Кремлёвская, 35)
(ч/н, 2-14 неделя) Дискретная математика. Пшеничный П.В., ауд. 1211 (Кремлёвская, 35)

12:10-13:40
Физкультура

13:50-15:20
(1-15 неделя) Лабораторный практикум по технологиям программирования. Шорина О.А., ауд. 809 (Кремлёвская, 35)`)
		
		case "чт":
			return c.Send(`📅 Четверг

10:10-11:40
(1-9 неделя) Технологии и методы программирования. Шорина О.А., ауд. 910 (Кремлёвская, 35)

12:10-13:40
(1-9 неделя) Документоведение. Рубцова Р.Г., ауд. 809 (Кремлёвская, 35)

13:50-15:20
(1-9 неделя) Основы управленческой деятельности. Овчинников М.Н., ауд. 112 (Кремлёвская, 16А)

15:50-17:20
(1-14 неделя) Электроника и схемотехника. Галиуллин А.А., ауд. 007 (Кремлёвская, 16А)`)
		
		case "пт":
			return c.Send("📅 Пятница\n\nЗанятий нет")
		
		case "сб":
			return c.Send(`📅 Суббота

8:30-10:00
(1-3 неделя) Операционные системы. Мубараков Б.Г., ауд. 907 (Кремлёвская, 35)
(1-15 неделя) Операционные системы. Мубараков Б.Г., ауд. 907 (Кремлёвская, 35)

12:10-13:40
Физкультура`)
		
		case "вс":
			return c.Send("📅 Воскресенье\n\nВыходной день")
		
		default:
			return c.Send("Неверный ввод. Пожалуйста, введите день недели (пн, вт, ср, чт, пт, сб, вс)")
		}
	})
}