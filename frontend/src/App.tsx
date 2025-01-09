import { Layout } from "antd";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Detail from "./pages/detail/cml";
import Infomation from "./pages/infomation";


export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Infomation />}> </Route>
        <Route path="/detail/:id" element={<Detail />} />
      </Routes>
    </BrowserRouter>
  );
}
