package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hibernateGolang/sk.kasv.huzvare/reviews"
	"hibernateGolang/sk.kasv.huzvare/structs"
	"hibernateGolang/sk.kasv.huzvare/testDB"
	"log"
)

var (
	addStudent = structs.Student{
		Email:     "marek@marek.sk",
		FirstName: "Marek",
		LastName:  "Les",
	}
	updateStudent = structs.Student{
		Email:     "ajsdljas",
		FirstName: "",
		LastName:  "",
	}
	addInstructor = structs.Instructor{
		FirstName: "Jano",
		LastName:  "Janocko",
		Email:     "p.janocko@marek.sk",
	}
	addInstructorWithDetails = structs.Instructor{
		FirstName: "Jano",
		LastName:  "Janocko",
		Email:     "p.janocko@marek.sk",
		InstructorDetail: structs.InstructorDetail{
			YoutubeChannel: "marekovoLetPlay",
			Hobby:          "Mikulas",
		},
	}
	addCourse = structs.Course{
		Title: "Marek",
	}
	addInstructorDetails = structs.InstructorDetail{
		YoutubeChannel: "Jano",
		Hobby:          "Jano",
	}
	updatedInstructor = structs.InstructorUpdate{
		FirstName: "Matus",
		LastName:  "Kovacz",
		Email:     "kovaz",
	}
	addReview = structs.Review{
		CourseID: 10,
		Comment:  "Najlepsi course na svete",
		Rating:   5,
	}
	updateReview = structs.Review{
		CourseID: 10,
		Comment:  "I changed my mind it was great",
		Rating:   5,
	}
)

func main() {

	testDB.IsJdbcRunning(structs.Dsn)
	db, err := gorm.Open(mysql.Open(structs.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	////first add student and delete
	//student.CreateStudent(db, &addStudent)
	//student.DeleteStudentByID(db, 5)
	//updateStudent.Email = "jano.jano@jano.sk"
	//student.UpdateStudentByID(db, 4, updateStudent)
	//student.AddStudentToCourse(db, 4, 12)
	//student.RemoveStudentFromCourse(db, 4, 13)
	//studentForPrint := student.GetStudentByID(db, 4)
	//fmt.Println(studentForPrint)
	//studentsForPrint, _ := student.SearchStudents(db, "alice", "", "m")
	//fmt.Println(studentsForPrint)
	//studentsForPrint, _ = student.SearchStudents(db, "", "Brown", "")
	//fmt.Println(studentsForPrint)
	//studentsForPrint, _ = student.SearchStudents(db, "", "", "alice.johnson@email.com")
	//fmt.Println(studentsForPrint)
	//cours, _ := student.GetCoursesForStudent(db, 4)
	//fmt.Println(cours)

	//instructor.DeleteInstructorById(db, 14)
	//instructor.AddInstructor(db, addInstructor)
	//instructor.AddInstructor(db, addInstructorWithDetails)
	//instructor.UpdateInstructor(db, 5, updatedInstructor)
	////instructor.UpdateInstructorDetail(db, 5, addInstructorDetails)
	//instruc, _ := instructor.FindInstructorByDetails(db, "Eva", "", "")
	//fmt.Println(instruc)
	//instruc, _ = instructor.FindInstructorByDetails(db, "", "Green", "")
	//fmt.Println(instruc)
	//instruc, _ = instructor.FindInstructorByDetails(db, "", "", "grace.hall@email.com")
	//fmt.Println(instruc)
	//cours, _ := course.GetCoursesByInstructorID(db, 4)
	//fmt.Println(cours)

	//course.AddCourseToInstructor(db, 6, addCourse)
	//course.DeleteCourseByID(db, 14)
	//course.UpdateCourseByID(db, 13, "Marek")
	//cour, _ := course.FindCoursesByTitle(db, "Mathematics 101")
	//fmt.Println(cour)

	//cour, _ := course.GetCoursesByInstructorID(db, 3)
	//fmt.Println(cour)
	reviews.AddReview(db, &addReview)
	//reviews.UpdateReviewByID(db, 5, updateReview)
	//result, _ := reviews.GetCourseWithReviews(db, 10)
	//fmt.Println(result)

	//cours, _ := instructor.ListCoursesByInstructorID(db, 4)
	//fmt.Println(cours)

	//cour, _ = student.GetCoursesForStudent(db, 1)
	//fmt.Println(cour)

	/* zhrnutie kodu structs su nieco ako Konstruktory v jave a tu je to tak ze ked dame bud vsetky
	infromacie alebo len jednu stale to berie a vyplni len tamtie a hlavne co GO neni uplne objektovo orientovany
	mam tam classy jak v jave ale neni to vobec treba a je tam vela z C cka takze "&" pred objektom odkazuje na to ze
	kde je v ramke a hlavne nejake metody su take iste ale su spravene inymi sposobmi vid funkcie (FindInstructorByDetails)
	&& (SearchStudents) to je taky zaklad ked si to chcete vyskusat tak spustite dbTry1.sql a napiste do prikazoveho riadku
	go mod tidy a malo by to fungovat*/

	/*Poziadavky na HIBERNATE projekt:
	-        SQL script – vygenerovanie databazy, vsetkych entit + relacie
	-        SQL script – sample data
	Maven – project
	TestJDBC – kontrola spojenia
	Student
	Prepojenie na entity Course M:N
	o   Pridanie a odtsranit course
	o    genetovat json file s informaciami o studentovi a jeho kurzoch
	               InstrucTorDetail
	               Instructor

	Prepojenie na InstructorDetail 1:1
	Prepojenie na Course 1:N
	Pridat a odobrat kurz
	Course
	Prepojenie na studentov M:N
	Prepojenie na intructor N:1
	Review
	Prepojenie N:1 s entitou Course
	Moznost pridavat kurzom review – comment + rating
	-        User stories:
	-        Vytvorit a zmazat studenta
	-        Zmenit studentovi osobne informacie
	-        Pridat alebo odstranit kurz ktory navstevuje
	-        Vygenerovat JSON o studentovi – vsetky info
	-        Zmena udajov – update
	-        Vyhladanie studentov – HQL – bud podla mena, alebo podla emailu
	-        Vypisanie kurzov ktore student ma priradene
	-
	-        Pridat a zmazat instruktora
	-        Pridat instruktorovi instruktordetail
	-        Pridat alebo odstranit kurz
	-        Zmena udajov
	-        Vypisat kurzy ktore instruktor uci
	-        Vyhladat instruktora – HQL
	-        Vytvorit, zmazat course
	-        Zmena kurzu
	-        Vyhladat kurz – HQL
	-        Pridat hodnotenie kurzu a rating
	-        Vypisat hodnotenia a ratingy vybraneho kurzu
	*/

}
