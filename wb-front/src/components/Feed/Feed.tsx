import React from "react";
import { Card } from "antd";
import styles from "./feed.module.scss";

const Feed: React.FC = () => {
  return (
    <div className={styles.feed}>
      <Card className={styles.post}>
        <h3>Post Title 1</h3>
        <p>Post content goes here...</p>
      </Card>
      <Card className={styles.post}>
        <h3>Post Title 2</h3>
        <p>Another post content...</p>
      </Card>
      <Card className={styles.post}>
        <h3>Post Title 3</h3>
        <p>Yet another post content...</p>
      </Card>
    </div>
  );
};

export default Feed;