const _connectToDevServerUpdates = () => {
    const socket = new WebSocket('ws://localhost:4000/ws');

    socket.onopen = () => {
        if (window._reconnectToDevServerInterval) {
            setTimeout(() => {
                window.location.href = window.location.href;
            }, 500);
        }

        clearInterval(window._reconnectToDevServerInterval);

        socket.onclose = () => {
            window._reconnectToDevServerInterval= setInterval(() => {
                _connectToDevServerUpdates();
            }, 5000);
        }
    }

    socket.onmessage = () => {
        window.location.href = window.location.href;
    }

    socket.onerror = () => {
        socket.close();
    }
}

_connectToDevServerUpdates();
