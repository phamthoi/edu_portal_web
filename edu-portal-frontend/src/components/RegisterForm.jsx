import {useState} from "react";

function RegisterForm({onSuccess}){
    const [form, setForm] = useState({
        username: "",
        password: "",
        fullName: "",
        email: "",
    });

    const handleChange = (e) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        try{
            const res = await fetch("http://localhost:8080/register",{
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify(form),
            });

            const data = await res.json();

            if (res.ok){
                alert("you have successfully logged in! please login.");
                onSuccess(); //return login page
            }else{
                alert(data.error ||"failed");
            }
        } catch(err){
            alert("Error when try connect with server");
        }
    };

    return(
        <form onSubmit={handleSubmit} className="bg-white p-6 rounded shadow-md max-w-sm mx-auto mt-10">
            <h2 className="text-xl font-semibold mb-4 text-center">Register</h2>

            <input
                name="username"
                placeholder="Username"
                className="w-full mb-3 border p-2 rounded"
                onChange={handleChange}
            />

            <input
                name="password"
                type="password"
                placeholder="Password"
                className="w-full mb-3 border p-2 rounded"
                onChange={handleChange}
            />

            <div className="w-full mb-3">
                <label htmlFor="role" className="block text-sm font-medium text-gray-700 mb-2">
                    Select your role
                </label>
                <select
                    id="role"
                    name="role"
                    className="block w-full p-2 border border-gray-300 rounded-lg shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
                    onChange={handleChange}
                >
                    <option value="">-- Choose a role --</option>
                    <option value="student">Student</option>
                    <option value="teacher">Teacher</option>
                    <option value="admin">Admin</option>
                </select>
            </div>

            <input
                name="fullName"
                placeholder="Full Name"
                className="w-full mb-3 border p-2 rounded"
                onChange={handleChange}
            />

            <input
                name="email"
                placeholder="Email"
                className="w-full mb-3 border p-2 rounded"
                onChange={handleChange}
            />

        <button className="bg-green-500 text-white px-4 py-2 rounded w-full">Register</button>
    </form>

    );
}

export default RegisterForm;