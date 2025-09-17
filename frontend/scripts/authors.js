const authorList = document.getElementById('authorList');

function getAuthors(url, endpoint, id) {
    const fullUrl = id ? `${url}/${endpoint}/${id}` : `${url}/${endpoint}`;

    return fetch(fullUrl)
        .then((response) => response.json())
        .then((data) => {

            data.forEach((singleData) => {
                authorList.innerHTML += `

                            <div class="col-lg-2 col-md-3 col-sm-4 col-6">
                                <div class="card bg-transparent border-0 text-white h-100">
                                    <div class="card-body text-center p-3">
                                        <div class="bg-dark rounded-circle mx-auto mb-3 d-flex align-items-center justify-content-center"
                                            style="width: 120px; height: 120px;">
                                            <svg width="60" height="60" viewBox="0 0 24 24" fill="none"
                                                xmlns="http://www.w3.org/2000/svg">
                                                <path
                                                    d="M12 12C14.7614 12 17 9.76142 17 7C17 4.23858 14.7614 2 12 2C9.23858 2 7 4.23858 7 7C7 9.76142 9.23858 12 12 12Z"
                                                    fill="#9CA3AF" />
                                                <path
                                                    d="M12 14C8.13401 14 5 17.134 5 21C5 21.5523 5.44772 22 6 22H18C18.5523 22 19 21.5523 19 21C19 17.134 15.866 14 12 14Z"
                                                    fill="#9CA3AF" />
                                            </svg>
                                        </div>
                                        <h6 class="card-title mb-1 fw-bold" style="color: black;">${singleData.name}</h6>
                                        <p class="card-text small text-muted mb-0">Author</p>
                                    </div>
                                </div>
                            </div>

                `;
            });
        })
        .catch((err) => console.log(err));
}

getAuthors("http://localhost:8080", "authors", null);