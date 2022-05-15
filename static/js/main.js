
    // let list = document.getElementByClassName("cat-list")
    
    showCategories()
    showExpenses()
    async function showCategories() {
        
        const responseCat = await fetch("http://localhost:4000/categories")
        const dataCat = await responseCat.json()
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
            li.innerHTML = "Expense name - " + item.expenseName + " " + "Category: " + item.expenseCat
                return li
            })
            expLists.forEach(list => {
                document.querySelector(".exp-list").appendChild(list)
            });
    }