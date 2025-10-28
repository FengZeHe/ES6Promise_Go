const axios = require('axios');

// function hi() {
//     axios.get("http://localhost:8080/hi").then((res) => {
//         return res.data;
//     }).then((data) => {
//         const strData = JSON.stringify(data);
//         rep(strData);
//     })
// }

function rep(msg) {
    console.log("收到回复" + msg)
}

// hi()

// async function hi() {
//     try {
//         const res = await axios.get("http://localhost:8080/hi");
//         const data = res.data;
//         const strData = JSON.stringify(data);
//         rep(strData);
//     } catch (err) {
//         console.log("ERROR:", err)
//     }
// }

// hi()

// setTimeout(()=>{
//     console.log("1s后打印")
// },1000)

// console.log("同步代码")


async function hi() {
    const res = await axios.get("http://localhost:8080/hi");
    return res.data;
}

async function callHi(times) {
    try {
        const list = Array.from({ length: times }, () => hi())
        const results = await Promise.all(list);

        console.log("结果是:", results)
    } catch (err) {
        console.log("错误:", err)
    }
}

callHi(2)

