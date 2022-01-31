const sendRequest = (method, url) => {
    const headers = {
        'Content-Type': 'application/json',
    };

    return fetch(url, {
        method: method,
        headers: headers
    }).then(async (response) => ({
        status: response.status,
Иван
Иван 16:50

        data: await response.json(),
    }))
};

document.querySelector('#vdovin-button').addEventListener('click', () => {
    sendRequest('GET', 'http://84.38.181.200:7171/api/vdovin').then(data => {
        alert('Здравствуйте, ' + data.data.first_name + ' ' + data.data.last_name);
    });
});
