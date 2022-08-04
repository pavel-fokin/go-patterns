function App() {
    return (
        <button
            onClick={() => { console.log("Clicked") }}
        >
            Like It
        </button>
    )
}

const root = ReactDOM.createRoot(document.getElementById("app"));
root.render(<App />);
