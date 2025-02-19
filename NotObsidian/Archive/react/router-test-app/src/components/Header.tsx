// import { Link } from 'react-router-dom'
import { NavLink } from 'react-router-dom'

const Header = () => {
  return (
    <div className="card bg-light">
        <div className="card-heading">
            <h2 className="text-center m-3">Foxes and Fossils</h2>
            <p>
                <a href="http://foxesandfossils.com">https://foxesandfossils.com</a>
            </p>
            <div className="row">
                <div className="col-12">
                    <NavLink to="/home" className={({ isActive }) => {
                        return isActive ? "btn menu btn-dark" : "btn menu btn-success";
                    }}>
                        Home
                    </NavLink>
                    <NavLink to="/about" className={({ isActive }) => {
                        return isActive ? "btn menu btn-dark" : "btn menu btn-success";
                    }}>
                        About
                    </NavLink>
                    <NavLink to="/members" className={({ isActive }) => {
                        return isActive ? "btn menu btn-dark" : "btn menu btn-success";
                    }}>
                        Members
                    </NavLink>
                    <NavLink to="/songs" className={({ isActive }) => {
                        return isActive ? "btn menu btn-dark" : "btn menu btn-success";
                    }}>
                        Songs
                    </NavLink>
                    {/* <Link className="btn btn-success menu" to="/">
                        Home
                    </Link>
                    <Link className="btn btn-success menu" to="/about">
                        About
                    </Link>
                    <Link className="btn btn-success menu" to="/members">
                        Members
                    </Link>
                    <Link className="btn btn-success menu" to="/songs">
                        Songs
                    </Link> */}
                </div>
            </div>
        </div>
    </div>
  )
}

export default Header;