import React, { Component } from 'react';
import './Message.css';

class Message extends Component {
  constructor(props) {
    super(props);
    let temp = JSON.parse(this.props.message);
    this.state = {
      message: temp
    }
  }
  
  render() {
    return (
      <div>
        <p>username:&nbsp; 
        {this.state.message.body}
        </p>
      </div>
    );
  };

}

export default Message;