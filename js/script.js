
getPrey()

function getPrey() {
    console.log("Hello world")
    const xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            document.getElementById("demo").innerHTML = this.responseText;
        }
    };
    xhttp.open("GET", "http://localhost:1234/v1/prey/2", true);
    xhttp.send();
}