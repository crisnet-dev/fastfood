const productsList = document.querySelector('.products-list');
const popup = document.getElementById('popup-container');

const websocket = new WebSocket("/ws?password=Hj$Xx33@");

function searchOrder() {
    const input = document.querySelector('');
}

document.addEventListener('keydown', (e) => {
    if (e.ctrlKey && e.key.toLowerCase() === 'q') {
        fetch('/order', {
            method: 'DELETE'
        })
            .then(response => response.json())
            .then(data => {
                if (data?.message) {
                    alert(data.message);
                }
            });
    }
});

function closePopup() {
    popup.style.display = 'none';
}

function listProducts(data) {
    const total = data.reduce((acc, product) => {
        return acc + (product.price * product.quantity);
    }, 0);

    document.getElementById('total-money').textContent = `Total ${total.toFixed(2)} KZ`;

    console.log(data);

    data.forEach(product => {
        const productItem = document.createElement('div');
        const span1 = document.createElement('span');
        const span2 = document.createElement('span');
        const span3 = document.createElement('span');

        productItem.classList.add('product-item');

        span1.textContent = product.product_name;
        span2.textContent = `Quantidade: ${product.quantity}`;
        span3.textContent = `${product.price} KZ`;

        productItem.appendChild(span1);
        productItem.appendChild(span2);
        productItem.appendChild(span3);

        productsList.appendChild(productItem);
    });
}

function listOrders(data) {
    const orderListContainer = document.querySelector('.order-list-container');

    orderListContainer.textContent = "";

    data.forEach((order) => {
        const btnDetails = document.createElement("button");
        const orderCard = document.createElement("div");
        const orderInfo = document.createElement("div");
        const avatar = document.createElement("div");
        const name = document.createElement("span");
        const time = document.createElement("span");

        name.textContent = order.name;
        time.textContent = order.time;
        btnDetails.textContent = "Ver Detalhes";

        orderCard.classList.add('order-card');
        orderInfo.classList.add('order-info');
        btnDetails.classList.add('btn-info');
        avatar.classList.add('avatar');

        orderInfo.appendChild(avatar);
        orderInfo.appendChild(name);
        orderInfo.appendChild(time);

        orderCard.appendChild(orderInfo);
        orderCard.appendChild(btnDetails);

        orderListContainer.appendChild(orderCard);

        btnDetails.addEventListener('click', () => {
            popup.style.display = 'flex';
            document.getElementById('clientName').textContent = `Nome: ${order.name}`;
            document.getElementById('clientLocation').textContent = `Localização: ${order.location}`;
            listProducts(order.products);
        });
    });
}

function addOrder(order) {
    const orderListContainer = document.querySelector('.order-list-container');

    const btnDetails = document.createElement("button");
    const orderCard = document.createElement("div");
    const orderInfo = document.createElement("div");
    const avatar = document.createElement("div");
    const name = document.createElement("span");
    const time = document.createElement("span");

    const now = new Date();

    name.textContent = order.name;
    time.textContent = order.time;
    btnDetails.textContent = "Ver Detalhes";

    orderCard.classList.add('order-card');
    orderInfo.classList.add('order-info');
    btnDetails.classList.add('btn-info');
    avatar.classList.add('avatar');

    orderInfo.appendChild(avatar);
    orderInfo.appendChild(name);
    orderInfo.appendChild(time);

    orderCard.appendChild(orderInfo);
    orderCard.appendChild(btnDetails);

    orderListContainer.appendChild(orderCard);

    btnDetails.addEventListener('click', () => {
        popup.style.display = 'flex';
        document.getElementById('clientName').textContent = `Nome: ${order.name}`;
        document.getElementById('clientLocation').textContent = `Localização: ${order.location}`;
        listProducts(order.products);
    });
}

websocket.onopen = () => {
    // alert("Connected");
    document.querySelector('#status').textContent = "Connectado";
}

websocket.onmessage = (data) => {
    const message = JSON.parse(data.data);

    console.log(message);

    if (message.type != "UpdateAdminCounter" || message.type != "Error") {
        // alert(`Novo pedido\nVeja no console`);

        if (message?.orders) {
            listOrders(message?.orders);
        } else if (message?.name) {
            addOrder(message);
        }

    }

    if (message?.type === "UpdateAdminCounter") {
        document.querySelector('#admin-counter').textContent = `${message.admin_counter} Admins`;
    }

    if (message?.type === 'Error') {
        alert(message.message);
    }

}

websocket.onclose = (err) => {
    document.querySelector('#status').textContent = "Não connectedo";
    console.log(err);
}
