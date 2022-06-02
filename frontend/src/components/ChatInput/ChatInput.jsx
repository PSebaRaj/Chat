import React, { Component } from 'react';
import './ChatInput.css';

class ChatInput extends Component {
  
  render() {
    return (
      <div>
        <input onKeyDown={this.props.send} placeholder="Type a message... Hit Enter to Send"/>
      </div>
    );
  };

}

export default ChatInput;