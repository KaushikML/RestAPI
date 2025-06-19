const API_URL = 'http://localhost:8080';
let token = '';

function showResult(id, data) {
    document.getElementById(id).textContent = JSON.stringify(data, null, 2);
}

document.getElementById('signupForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const email = document.getElementById('signupEmail').value;
    const password = document.getElementById('signupPassword').value;

    const res = await fetch(`${API_URL}/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
    });
    showResult('signupResult', await res.json());
});

document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const email = document.getElementById('loginEmail').value;
    const password = document.getElementById('loginPassword').value;

    const res = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
    });
    const data = await res.json();
    if (res.ok && data.token) {
        token = data.token;
    }
    showResult('loginResult', data);
});


document.getElementById('loadEvents').addEventListener('click', async () => {
    const res = await fetch(`${API_URL}/events`);
    showResult('eventsResult', await res.json());
});

document.getElementById('updateEventForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const id = document.getElementById('updateEventId').value;
    const name = document.getElementById('updateEventName').value;
    const description = document.getElementById('updateEventDescription').value;
    const location = document.getElementById('updateEventLocation').value;
    const rawDate = document.getElementById('eventDateTime').value;
    const dateTime = rawDate ? `${rawDate}:00Z` : '';

    const res = await fetch(`${API_URL}/events/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        },
        body: JSON.stringify({ name, description, location, dateTime })
    });
    showResult('updateEventResult', await res.json());
});

document.getElementById('deleteEventForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const id = document.getElementById('deleteEventId').value;
    const res = await fetch(`${API_URL}/events/${id}`, {
        method: 'DELETE',
        headers: { 'Authorization': token }
    });
    showResult('deleteEventResult', await res.json());
});