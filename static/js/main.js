
    // let list = document.getElementByClassName("cat-list")
    showExpenses()
    showCategories()
    showCategoriesFilter()

    let filter = document.querySelector('#filter')
    filter.addEventListener("click", getExpensesByCat)

    async function getExpensesByCat(){
        const check = document.getElementsByClassName("filter-cat").checked
        // const c = check.filter(x => {x.checked})

        console.log(check)
    }
   
    async function showCategoriesFilter(){
        const dataCat = await getCategories() 
        let catLists = dataCat.map(item => {
            let li = document.createElement("li")
            li.innerHTML = "<input class='filter-cat' name='expenseCat' type='checkbox' value='" + item.catName +"'>"+"<label for='catname'>"+item.catName+"</label>"
                return li
            })
            catLists.forEach(list => {
                document.querySelector(".cat-list-filter").appendChild(list)
            });
    }

    async function getCategories(){
        const responseCat = await fetch("http://localhost:4000/categories")
        const dataCat = await responseCat.json()
        return dataCat
    }
    async function showCategories() {
        const dataCat = await getCategories()
        console.log(dataCat)
        let catLists = dataCat.map(item => {
            let li = document.createElement("li")
            li.innerHTML = "<input name='expenseCat' type='radio' value='" + item.catName +"'>"+"<label for='catname'>"+item.catName+"</label>"
                return li
            })
            catLists.forEach(list => {
                document.querySelector(".cat-list").appendChild(list)
            });
    }

    async function showExpenses(){
        const responseExp = await fetch("http://localhost:4000/expenses")
        const dataExp = await  responseExp.json()
        console.log(dataExp)
        let expLists = dataExp.map(item=>{
            let li = document.createElement("li")
            li.classList.add("recent-exp-li")
            li.innerHTML = "Expense name - " + item.expenseName + " " + "Category: " + item.expenseCat + " " + "DATE:" + item.expenseDate + "<button class='del-button' onClick=deleteExpense(this.value) value='"+ item._id+"'> Delete </button>"
                return li
            })
            expLists.forEach(list => {
                document.querySelector(".exp-list").appendChild(list)
            });
    }

    async function deleteExpense(ExpenseId){
        const responseExp = await fetch("http://localhost:4000/expenses/"+ExpenseId, {
            method: "DELETE",
        })
        const ul = document.querySelector(".exp-list")
        ul.innerHTML = ""
        showExpenses()
    }

    document.querySelectorAll('.del-button').forEach(button => {
        console.log(button.value)
        button.addEventListener('click', () => {
            const expenseId = this.value
            console.log("cliaced")
            deleteExpense(expenseId)
        });
    });
    