import React from "react";
import Header from "../../components/Header/Header";
import MainContent from "../../components/Layout/Layout";
import styles from "./home.module.scss"

const Home: React.FC = () =>{
  return (
    <div className={styles.mainContainer}>
      <Header />
      <MainContent />
    </div>
  );
}

export default Home;