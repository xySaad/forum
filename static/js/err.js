export function createPopup(err) {
    let div = document.createElement('div');
    div.style.position = 'fixed';
    div.style.top = '20px';
    div.style.right = '20px';
    div.style.padding = '15px';
    div.style.background = 'red';
    div.style.color = 'white';
    div.style.borderRadius = '5px';
    div.style.boxShadow = '0 4px 6px rgba(0,0,0,0.1)';
    div.style.zIndex = '1000';
    div.style.fontSize = '16px';
    div.style.fontWeight = 'bold';
    div.textContent = err;
    document.body.appendChild(div);
    setTimeout(() => {
        div.remove();
    }, 3000);
}
