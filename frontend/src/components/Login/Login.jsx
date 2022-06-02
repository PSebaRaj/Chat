import React, { Component } from 'react';
import './Login.css';

class Login extends Component {
    // state = {
    //     username: '',
    //     password: '',
    // }

    // handleChange = (event) => {
    //     const {username, password} = event.target;
    //     this.setState({[username]:password});
    // }

    // handleSubmit = (event) => {
    //     // event.preventDefault(); // I WANT IT TO REFRESH UPON LOGIN
    //     this.props.login(this.state.username, this.state.password);
    // }


    // render() {
    //   return (
    //     <div>
    //         <form onSubmit = {this.handleSubmit}>
    //             <input type="username" placeholder="Username..." required onChange={this.handleChange}/>
    //             <input type="password" placeholder="Password..." required onChange={this.handleChange}/>
    //             <button type="submit" onSubmit={this.handleSubmit}>Login</button>
    //         </form>
    //     </div>
    //   );
    // };
    render() {
        return (
          <div>
            <input onKeyDown={this.props.send} placeholder="Display Name..."/>
          </div>
        );
      };
  }
  
  export default Login;