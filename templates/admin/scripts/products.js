const formData = new FormData();

const fileInput = document.getElementById("file-input");
const inputProductName = document.getElementById("input-product-name");
const inputProductPrice = document.getElementById("input-product-price");

const loadImage = () => {
    const file = fileInput.files[0];

    const reader = new FileReader();

    reader.onload = () => {
        const base64 = reader.result;
        formData.append("image", file);
        document.getElementById('img-preview').src = base64;
    }

    reader.readAsDataURL(file);
}

const sendDataToServer = async (formData) => {
    document.querySelector('.btn-add-product-text').style.display = 'none';
    document.querySelector('.loading').style.display = 'block';

    const response = await fetch("/product/upload", {
        method: "POST",
        body: formData
    });
    const data = response => response.json();

    document.querySelector('.btn-add-product-text').style.display = 'block';
    document.querySelector('.loading').style.display = 'none';

    console.log(data);
    

    if (response.status !== 200) {
        alert("Não foi possível realizar a operação de momento!\n Para mais detalhes veja o console.");
        console.log(data);
        return;
    }

    alert("Sucesso!");

    inputProductName.value = "";
    inputProductPrice.value = "";
    document.getElementById('img-preview').src = "";
}

document.querySelector('button').addEventListener('click', () => {

    if (inputProductName.value === "" || inputProductPrice.value === " " || !fileInput.files) {
        alert("Digite os dados requisitados!");
        return;
    }

    formData.append(
        "data",
        JSON.stringify({
            id: 0,
            product_name: inputProductName.value,
            price: parseFloat(inputProductPrice.value),
            image_url: ""
        })
    );

    sendDataToServer(formData);
});

fileInput.addEventListener('change', loadImage);
