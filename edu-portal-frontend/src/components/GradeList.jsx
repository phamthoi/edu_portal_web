import { useEffect, useState } from "react";

function GradeList(){
    const [grades, setGrades] = useState([]);
    const [loading, setLoading] = useState(true);
    useEffect(() =>{
        const token = localStorage.getItem("token");

        fetch("http://localhost:8080/")
    })
}