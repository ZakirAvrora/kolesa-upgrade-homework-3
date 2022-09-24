<?php
namespace App\Model;

class Cat {
   protected string $id;
   protected  string $imageUrl;

   public function __construct(string $id, string $imageUrl)
   {
        $this->id = $id;
        $this->imageUrl = $imageUrl;
   }

   public function getId(): string {
        return $this->id;
   }

   public function getImage(): string{
        return $this->imageUrl;
   }
}