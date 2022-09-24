<?php
//index.php
$autoloadPath = __DIR__ . '/vendor/autoload.php';

require_once $autoloadPath;

use App\Model\Contoller;
use App\Model\Router;

$router = new Router();

$router->get('/', function(){
    $controller = new Contoller();
    $Breeds = $controller->requestBreeds();

    require_once __DIR__.'/static/templates/homePage.phtml';
});

$router->get('/cats', function(){
    $limit = $_GET['limit'];
    $breed = $_GET['breed'];

    if(!isset($limit) || !isset(  $breed) ) {
        header("HTTP/1.0 400 Bad Request");
        require_once __DIR__.'/static/templates/400.phtml';
        return;
    }

    $controller = new Contoller();
    $Cats = $controller->requestCats($limit,$breed);
    
    if (empty($Cats)) {
        header("HTTP/1.0 400 Bad Request");
        require_once __DIR__.'/static/templates/400.phtml';
        return;
    }

    require_once __DIR__.'/static/templates/search.phtml';
});

$router->get('/details', function(){
    $id = $_GET['id'];
    $controller = new Contoller();
    $Cat = $controller->requestDetails($id);
    if(!isset($Cat)){
        header("HTTP/1.0 400 Bad Request");
        require_once __DIR__.'/static/templates/400.phtml';
        return;
    }

    require_once __DIR__.'/static/templates/details.phtml';
});

$router->addNotFoundHandler(function(){
    require_once __DIR__.'/static/templates/404.phtml';
});

$router->run();

?>

