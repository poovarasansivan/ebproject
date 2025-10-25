import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Header from "./components/header";
import Home from "./pages/home";

function App() {
  return (
    <Router>
      <div className="flex-1 flex flex-col">
        <Header />
        <div className="p-4 flex-1">
          <Routes>
            {/* <Route path="/" element={<Navigate to="/dashboard" />} /> */}
            <Route path="/home" element={<Home />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;
