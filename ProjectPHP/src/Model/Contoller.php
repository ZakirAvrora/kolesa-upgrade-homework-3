<?php

namespace App\Model;

use App\Model\Cat;
use App\Model\CatDetails;
use  GuzzleHttp\Client;
use GuzzleHttp\Exception\ClientException;

// define('BREED_LIST_API', 'https://api.thecatapi.com/v1/breeds');

class Contoller
{

    private Client $client;
     
    public function __construct()
    {
        $this->client = new Client([
            'base_uri' => 'https://api.thecatapi.com',
            'timeout'  => 2.0,
        ]);
    }

    public function requestBreeds() : array {
        $response = $this->client->request('GET','/v1/breeds/');
        $body = $response->getBody();
        $arr_breeds = json_decode($body);
        
        $breeds = array();

        foreach($arr_breeds as $breed){
            $breeds[$breed->name] = $breed->id;
        }

        return $breeds;
    }

    public function requestCats(string $limit, string $breed): array {
        $Cats = array();

        $response = $this->client->request('GET','/v1/images/search',[
            'query' =>[
                'limit'=>$limit,
                'breed_ids'=>$breed,
            ]
        ]);

        $body = $response->getBody();
        $arr_body = json_decode($body, true);
        
        foreach($arr_body as $obj){
            $cat = new Cat($obj['id'], $obj['url']);
            array_push($Cats, $cat);
        }

        return $Cats;
    }


    public function requestDetails(string $id) {
       
       try {
            $response = $this->client->request('GET','/v1/images/'.$id);
       }catch(ClientException $e){
            return null;
       }

        $body = $response->getBody();
        $obj = json_decode($body);
        
        $id = $obj->id;
        $url = $obj->url;

        if (!isset($obj->breeds)){
            $weight = "";
            $name = "";
            $temperament = "";
            $origin = "";
            $description = "";
            $lifespan = "";
            $extraInfoExist = false;
        }else{
            foreach($obj->breeds as $detail) {
                $weight = $detail->weight->metric;
                $name = $detail->name;
                $temperament = $detail->temperament;
                $origin = $detail->origin;
                $description = $detail->description ;
                $lifespan = $detail->life_span;
                $extraInfoExist = true;
            }
        }
        
        $Cat = new CatDetails($id, $url,$weight,$name,$origin,$temperament,$description,$lifespan);
        $Cat->isExtraInfo($extraInfoExist);

        return $Cat;
    }
}