import React, { Component } from "react";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import Login from "./components/Login/Login";
import "./App.css";
import { connect, sendMsg } from "./api";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "Anonymous",
      chatHistory: [],
    };

    this.send = this.send.bind(this);
    this.login = this.login.bind(this);
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message");
      this.setState((prevState) => ({
        chatHistory: [...prevState.chatHistory, msg],
      }));
      console.log(this.state);
    });
  }

  send(event) {
    // Enter
    if (event.keyCode === 13) {
      sendMsg(this.state.username + ": " + event.target.value);
      event.target.value = "";
    }

    // // Right arrow
    // if (event.keyCode === 39) {
    //   const usr = event.target.value;
    //   this.setState(() => ({
    //     username: usr,
    //   }));
    //   event.target.value = "";
    // }
  }

  login(event) {
    // Right arrow
    if (event.keyCode === 39) {
      const usr = event.target.value;
      this.setState(() => ({
        username: usr,
      }));
      event.target.value = "";
    }
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
        <Login send={this.login} />
      </div>
    );
  }
}

export default App;
