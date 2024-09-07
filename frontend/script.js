const URL = "http://localhost:8585";

const updateTokens = async (cb) => {
    try {
        const resp = await fetch(`${URL}/api/tokens/update`, {
            method: "POST",
            credentials: "include",
            body: JSON.stringify({
                login: localStorage.getItem("login"),
            }),
        });
    
        const data = await resp.json();
        
        if(data.status === "success") {
            localStorage.setItem("accessToken", data.data.accessToken);
            return cb();
        };
    } catch(e) {
        console.log(e);
    };
};

const registration = async () => {
    try {
        const resp = await fetch(`${URL}/api/users/registration`, {
            method: "POST",
            credentials: "include",
            body: JSON.stringify({
                login: document.querySelector("#reg_login").value,
                password: document.querySelector("#reg_password").value,
                isAdmin: true,
            }),
        });
    
        const data = await resp.json();
        console.log(data);   
    } catch(e) {
        console.log(e);
    };
};

const authorization = async () => {
    try {
        const login = document.querySelector("#auth_login").value;

        const resp = await fetch(`${URL}/api/users/authorization`, {
            method: "POST",
            credentials: "include",
            body: JSON.stringify({
                login: login,
                password: document.querySelector("#auth_password").value,
            }),
        });
    
        const data = await resp.json();
        console.log(data);

        localStorage.setItem("login", login);
        localStorage.setItem("accessToken", data.data.accessToken);
    } catch(e) {
        console.log(e);
    };
};

const addFile = async (file, fileName) => {
    try {
        const resp = await fetch(`${URL}/api/images/add`, {
            method: "POST",
            credentials: "include",
            headers: {
                "Authorization": `Bearer ${localStorage.getItem("accessToken")}`,
            },
            body: JSON.stringify({
                login: localStorage.getItem("login"),
                authorName: document.querySelector("#file_authorName").value,
                fileName: fileName,
                file: file,
                title: document.querySelector("#file_title").value,
                description: document.querySelector("#file_description").value,
                createdAt: new Date(),
            }),
        });

        if(resp.status === 401) return updateTokens(() => addFile(file, fileName));
    
        const data = await resp.json();
        console.log(data);       
    } catch(e) {
        console.log(e);
    };
};

const deleteFile = async (file) => {
    try {
        const resp = await fetch(`${URL}/api/images/${file.id}`, {
            method: "DELETE",
            credentials: "include",
            headers: {
                "Authorization": `Bearer ${localStorage.getItem("accessToken")}`,
            },
        });

        if(resp.status === 401) return updateTokens(() => deleteFile(file));
    
        const data = await resp.json();
        console.log(data);        
    } catch(e) {
        console.log(e);
    };
}; 

const getAllFiles = async () => {
    try {
        const resp = await fetch(`${URL}/api/images/all`, {
            method: "GET",
            credentials: "include",
        });

        if(resp.status === 401) return updateTokens(() => getAllFiles());
    
        const data = await resp.json();
        const images = data.data.images;
        
        for(const el of images) {
            document.querySelector("body").innerHTML += `
                <h1>${el.imageTitle}</h1> 
                <h2>${el.imageDescription}</h2> 
                <h3>${el.imageId}</h3> 
                <img 
                    id="${el.imageId}"
                    src="${URL}/api/images/${el.imageId}"
                >
            `;
        };

        for(const el of document.querySelectorAll("img")) el.addEventListener("click", (e) => deleteFile(e.target));
    } catch(e) {
        console.log(e);
    };
};

const createAuthor = async () => {
    try {
        const avatarFile = document.querySelector("#create_author_file");
        const file = avatarFile.files[0]; 
        const fileName = file.name;
    
        const reader = new FileReader();
    
        reader.onload = async (event) => {
            const arrayBuffer = event.target.result;
            const formData = new Uint8Array(arrayBuffer);
    
            const resp = await fetch(`${URL}/api/authors/add`, {
                method: "POST",
                credentials: "include",
                headers: {
                    "Authorization": `Bearer ${localStorage.getItem("accessToken")}`,
                },
                body: JSON.stringify({
                    login: localStorage.getItem("login"), 
                    avatarName: fileName,
                    avatarData: Array.from(formData),
                    name: document.querySelector("#create_author_name").value, 
                    description: document.querySelector("#create_author_description").value,
                }),
            });
            
            if(resp.status === 401) return updateTokens(() => createAuthor());
    
            const data = await resp.json();
            console.log(data);
        };
    
        reader.readAsArrayBuffer(file);
    } catch(e) {
        console.error(e);
    };
};

const updateAuthor = async () => {
    try {
        const resp = await fetch(`${URL}/api/authors/${document.querySelector("#update_author_id").value}`, {
            method: "PUT",
            credentials: "include",
            headers: {
                "Authorization": `Bearer ${localStorage.getItem("accessToken")}`,
            },
            body: JSON.stringify({
                name: document.querySelector("#update_author_name").value, 
                description: document.querySelector("#update_author_description").value,
            }),
        });
        
        if(resp.status === 401) return updateTokens(() => updateAuthor());

        const data = await resp.json();
        console.log(data);
    } catch(e) {
        console.error(e);
    };
};

const deleteAuthor = async (id) => {
    try {
        const resp = await fetch(`${URL}/api/authors/${id}`, {
            method: "DELETE",
            credentials: "include",
            headers: {
                "Authorization": `Bearer ${localStorage.getItem("accessToken")}`,
            },
        });
        
        if(resp.status === 401) return updateTokens(() => deleteAuthor(id));

        const data = await resp.json();
        console.log(data);
    } catch(e) {
        console.error(e);
    };
};

const getAuthor = async (id) => {
    try {
        const resp = await fetch(`${URL}/api/authors/${id}`, {
            method: "GET",
            credentials: "include",
            headers: {
                "Authorization": `Bearer ${localStorage.getItem("accessToken")}`,
            },
        });
        
        if(resp.status === 401) return updateTokens(() => getAuthor());

        const data = await resp.json();
        const author = data.data.author;

        console.log(data);
        alert(`
            authorId: ${author.authorId} \n
            authorName: ${author.authorName} \n
            authorDescription: ${author.authorDescription} \n
        `);
    } catch(e) {
        console.error(e);
    };
};

const getAuthorImage = async (id) => {
    try {
        const resp = await fetch(`${URL}/api/images/author/${id}`, {
            method: "GET",
            credentials: "include",
        });

        if(resp.status === 401) return updateTokens(() => getAuthorImage());
    
        const data = await resp.json();
        const images = data.data.images;
        
        let block;

        for(const el of document.querySelectorAll(".images_authorBtn")) {
            if(el.id !== id) continue;
            block = el;
        };

        block.innerHTML = "";

        for(const el of images) {
            block.innerHTML += `
                <h1>${el.imageTitle}</h1> 
                <h2>${el.imageDescription}</h2> 
                <h3>${el.imageId}</h3> 
                <img 
                    id="${el.imageId}"
                    src="${URL}/api/images/${el.imageId}"
                >
            `;
        };

        for(const el of document.querySelectorAll("img")) el.addEventListener("click", (e) => deleteFile(e.target));
    } catch(e) {
        console.log(e);
    };    
}; 

const updateAvatarAuthor = (e, name) => {
    try {
        const file = e.target.files[0]; 
        const fileName = file.name;
    
        const reader = new FileReader();
    
        reader.onload = async (event) => {
            const arrayBuffer = event.target.result;
            const formData = new Uint8Array(arrayBuffer);
    
            const resp = await fetch(`${URL}/api/authors/avatar`, {
                method: "PUT",
                credentials: "include",
                headers: {
                    "Authorization": `Bearer ${localStorage.getItem("accessToken")}`,
                },
                body: JSON.stringify({
                    name: name,
                    avatarName: fileName,
                    avatarData: Array.from(formData),
                }),
            });
            
            if(resp.status === 401) return updateTokens(() => updateAvatarAuthor(e, name));
    
            const data = await resp.json();
            console.log(data);
        };
    
        reader.readAsArrayBuffer(file);
    } catch(e) {
        console.error(e);
    };    
};

const getAllAuthors = async () => {
    try {
        const resp = await fetch(`${URL}/api/authors/all`, {
            method: "GET",
            credentials: "include",
            headers: {
                "Authorization": `Bearer ${localStorage.getItem("accessToken")}`,
            },
        });

        const data = await resp.json();
        const authors = data.data.authors;

        const block = document.querySelector("#all_authors");
        block.innerHTML = "";
        
        for(const author of authors) {
            block.innerHTML += `
                <img
                    src="${URL}/api/authors/avatar/${author.authorName}"
                    alt="Author avatar"   
                    style="width: 100px; height: 100px;"
                > 
                <input type="file" id="${author.authorName}" placeholder="update_avatar_author" class="update_avatar_author">
                <button id="${author.authorId}" class="get_authorBtn">${author.authorName}</button>
                <button id="${author.authorId}" class="get_author_imageBtn">get</button>
                <button id="${author.authorId}" class="delete_authorBtn">Delete</button>
                <div id="${author.authorId}" class="images_authorBtn"></div>
                <p>---<p> 
            `;
        };

        for(const el of document.querySelectorAll(".update_avatar_author")) {
            el.addEventListener("change", (e) => updateAvatarAuthor(e, e.target.id));
        };
        for(const el of document.querySelectorAll(".get_authorBtn")) {
            el.addEventListener("click", (e) => getAuthor(e.target.id));
        };
        for(const el of document.querySelectorAll(".delete_authorBtn")) {
            el.addEventListener("click", (e) => deleteAuthor(e.target.id));
        };
        for(const el of document.querySelectorAll(".get_author_imageBtn")) {
            el.addEventListener("click", (e) => getAuthorImage(e.target.id));
        };
    } catch(e) {
        console.error(e);
    };    
};

const updateDataOfImage = async () => {
    try {
        alert("updateDataOfImage");
        const resp = await fetch(`${URL}/api/images/${document.querySelector("#update_data_of_image_id").value}`, {
            method: "PUT",
            credentials: "include",
            headers: {
                "Authorization": `Bearer ${localStorage.getItem("accessToken")}`,
            },
            body: JSON.stringify({
                title: document.querySelector("#update_data_of_image_title").value, 
                description: document.querySelector("#update_data_of_image_description").value,
            }),
        });
        
        if(resp.status === 401) return updateTokens(() => updateDataOfImage());

        const data = await resp.json();
        console.log(data);
    } catch(e) {
        console.error(e);
    };    
};

document    
    .querySelector("#sendFile")
    .addEventListener("change", (e) => {
        const file = e.target.files[0]; 
        const fileName = file.name;
    
        const reader = new FileReader();
    
        reader.onload = (event) => {
            const arrayBuffer = event.target.result;
            const formData = new Uint8Array(arrayBuffer);
    
            addFile(
                Array.from(formData), 
                fileName,
            );
        };
    
        reader.readAsArrayBuffer(file);
    });

document
    .querySelector("#all")
    .addEventListener("click", () => getAllFiles());

document
    .querySelector("#regBtn")
    .addEventListener("click", () => registration());  

document
    .querySelector("#authBtn")
    .addEventListener("click", () => authorization());    

document
    .querySelector("#update_data_of_image_btn")
    .addEventListener("click", () => updateDataOfImage());    

document
    .querySelector("#create_authorBtn")
    .addEventListener("click", () => createAuthor());

document
    .querySelector("#update_authorBtn")
    .addEventListener("click", () => updateAuthor());

document
    .querySelector("#all_authorsBtn")
    .addEventListener("click", () => getAllAuthors());

