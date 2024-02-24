document.addEventListener('DOMContentLoaded', () => {
    console.log('DOM полностью загружен и разобран');

    // Обработка отправки формы
    const form = document.querySelector('form');
    form.addEventListener('submit', handleFormSubmit);

    // Переключение темы
    const themeToggleBtn = document.querySelector('#theme-toggle');
    themeToggleBtn.addEventListener('click', toggleTheme);
});

function handleFormSubmit(event) {
    event.preventDefault(); // Предотвратить стандартную отправку формы
    const formData = new FormData(event.target);
    const data = Object.fromEntries(formData.entries());
    console.log('Данные формы:', data);
    // Здесь может быть код для отправки данных формы на сервер
}

async function fetchData(url = '') {
    try {
        const response = await fetch(url);
        if (!response.ok) throw new Error(`Ошибка HTTP: ${response.status}`);
        const data = await response.json();
        console.log(data);
        return data;
    } catch (error) {
        console.error('Ошибка при получении данных:', error);
    }
}

function toggleTheme() {
    const body = document.body;
    body.classList.toggle('dark-theme'); // предполагается, что класс 'dark-theme' определен в CSS
}

function manipulateDOM() {
    const element = document.querySelector('#some-element');
    // Добавление, удаление или изменение контента элемента
    element.textContent = 'Новый текст';
}