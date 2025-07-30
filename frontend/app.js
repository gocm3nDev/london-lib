fetch('http://localhost:8080/books').then(r => r.json()) // Fetch the book list from the API
    .then(data => {
        const tbody = document.getElementById('bookList');
        tbody.innerHTML = '';

        data.forEach(book => {
            const tr = document.createElement('tr');

            tr.innerHTML = `
                <td>${book.id}</td>
                <td>${book.bookName}</td>
                <td>${book.author}</td>
                <td>${book.quantity}</td>
                <td>${book.description}</td>
                <td>${book.published}</td>  
                <td>${book.page}</td>
                <td>${book.bookLanguage}</td>
            `;
            tbody.appendChild(tr);
        });
    }
    ).catch(e => { document.getElementById('bookList').textContent = 'Hata: ' + e.message; });
