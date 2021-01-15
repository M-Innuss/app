// DB SETUP

var DB *sql.DB

func InitDB() {
	connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", "user", "pass", "database")

	var err error
	DB, err = sql.Open("postgres", connstring)
	if err != nil {
		revel.INFO.Println("DB Error", err)
	}
	revel.INFO.Println("DB Connected")
}

func init() {

	revel.Filters = []revel.Filter{
		revel.PanicFilter,    // Recover from panics and display an error page instead.
		revel.CompressFilter, // Compress the result.
		revel.ActionInvoker,  // Invoke the action.
	}

	revel.OnAppStart(InitDB)

}