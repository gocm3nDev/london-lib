const table = document.getElementById('bookList');

function getData(url, endpoint, id) {
    const fullUrl = id ? `${url}/${endpoint}/${id}` : `${url}/${endpoint}`;

    return fetch(fullUrl)
        .then((response) => response.json())
        .then((data) => {

            data.forEach((singleData) => {
                table.innerHTML += `
                <tr>
                    <td>${singleData.id}</td>
                    <td>${singleData.bookName}</td>
                    <td>${singleData.author}</td>
                    <td>${singleData.quantity}</td>
                    <td>${singleData.description}</td>
                    <td>${singleData.published}</td>
                    <td>${singleData.page}</td>
                    <td>${singleData.bookLanguage}</td>
                    <td>
                        <button class="btn btn-sm btn-warning me-1" onclick="editBook('http://localhost:8080', 'books', ${singleData.id})"><i class="fas fa-edit"></i> Edit</button>
                    </td>
                    <td>
                        <button class="btn btn-sm btn-danger" onclick="removeBook('http://localhost:8080', 'books', ${singleData.id})"><i class="fas fa-trash"></i> Remove</button>
                    </td>
                </tr>
                `;
            });
        })
        .catch((err) => console.log(err));
}

function removeBook(url, endpoint, id) {
    const fullUrl = `${url}/${endpoint}/${id}`;

    fetch(fullUrl, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        }
    })
        .then(res => {
            if (!res.ok) throw new Error("An error occured");
            return res.json();
        })
        .then(data => console.log("Deleted:", data))
        .catch(err => console.error("Error:", err));
}

getData("http://localhost:8080", "books", null);