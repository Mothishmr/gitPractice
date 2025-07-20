const pnameBtn = document.getElementById("pname");
const pname = pnameBtn.value;

async function fetchData(){

    try{
        const response = await fetch(`https://pokeapi.co/api/v2/pokemon/${pnameBtn.value}`);
        
        if(!response.ok){
            throw new Error("check pokemon name");
        }
        
        const responseData = await response.json();
        console.log(responseData.name);
        console.log(responseData.weight);
        
        const pimg = responseData.sprites.front_default;
        const imgEle = document.getElementById("pimg");
        imgEle.src = pimg;
        imgEle.style.display = "block"

    }
    catch(error){
        console.error(error)
    }

}

