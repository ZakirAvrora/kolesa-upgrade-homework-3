<?php
namespace App\Model;

class CatDetails extends Cat {
    protected  string $weight;
    protected  string $name;
    protected  string $origin;
    protected  string $temperament;
    protected  string $description;
    protected  string $lifespan;
    protected bool $extraInfoExist;

    public function __construct(string $id, string $imageUrl, string $weight, 
    string $name, string $origin, string $temperament, 
    string $description, string $lifespan)
    {
        $this->id = $id;
        $this->imageUrl = $imageUrl;
        $this->weight = $weight;
        $this->name = $name;
        $this->origin = $origin;
        $this->temperament = $temperament;
        $this->description = $description;
        $this->lifespan = $lifespan;
    }

    public function isExtraInfo(bool $extraInfoExist){
       $this->extraInfoExist = $extraInfoExist;
    }

    public function getExtraInfo(): bool{
        return $this->extraInfoExist ;
    }
    public function getWeight(): string{
        return $this->weight;
    }

    public function getName(): string{
        return $this->name;
    }

    public function getOrigin(): string{
        return $this->origin;
    }

    public function getDescription(): string{
        return $this->description;
    }

    public function getLifespan(): string{
        return $this->lifespan;
    }


    public function getTemperament(): string{
        return $this->temperament;
    }
} 