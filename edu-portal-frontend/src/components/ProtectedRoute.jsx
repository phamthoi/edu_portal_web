import {Navigate} from "react-router-dom";

function ProtectedRoute({children}){
    const token = localStorage.getItem("token");

    if(!token){
        //if it doesn't token -> redirect onto login
        return <Navigate to ="/" replace/>;
    }

    //if have token -> render content(children)
    return children;
}

export default ProtectedRoute;

//navigate: it is component of react router to redirect
//replace: it help not hold page /dashboard in history
