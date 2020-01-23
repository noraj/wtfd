import Leaderboard from './leaderboard';
import AdminPage from './admin';
import MainPage from './main';
import LoginRegister from './login-register';

(function(){

console.log("test");

new LoginRegister();

switch(location.pathname){
  case '/admin': {
    new AdminPage();
    new Leaderboard();
    break;
  }
  case '/leaderboard': {
    new Leaderboard();
    break;
  }
  default: {
    new MainPage();
    import './dragscroll';
    break;
  }
}

})();
