package internals

import "github.com/hajimehoshi/ebiten/v2"

type Player struct{
	Sprite *ebiten.Image
	XCoordinate float64
	YCoordinate float64
	LeftLaser *Laser
	RightLaser *Laser
	// Scale float64
}


func (p *Player) move(speed float64){
	if ebiten.IsKeyPressed(ebiten.KeyW){
		p.YCoordinate -= speed
		
	}
	if ebiten.IsKeyPressed(ebiten.KeyS){
		p.YCoordinate += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.XCoordinate -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.XCoordinate += speed
	}
	
}

func (p *Player) clamp_player(windowW float64, windowH float64){
	if p.YCoordinate < 0 {
		p.YCoordinate = 0
	}

	if p.YCoordinate>windowH - float64(p.Sprite.Bounds().Dy()){
		p.YCoordinate = windowH - float64(p.Sprite.Bounds().Dy())
	}

	if p.XCoordinate < 0 {
		p.XCoordinate = 0
	}
	if p.XCoordinate > windowW- float64(p.Sprite.Bounds().Dx()) {
		p.XCoordinate = windowW - float64(p.Sprite.Bounds().Dx())
	}
}

func (p *Player) shoot(speed float64){
	if p.LeftLaser != nil && p.RightLaser!=nil{
		p.LeftLaser.Move(speed)
		p.RightLaser.Move(speed)
	}
}

func (p *Player) generateLaser(){
	p.LeftLaser = &Laser{Sprite:MustLoadImage(LASER_BLUE_16), XCoordinate: p.XCoordinate, YCoordinate: p.YCoordinate}
	p.RightLaser = &Laser{Sprite:MustLoadImage(LASER_BLUE_16), XCoordinate: p.XCoordinate + float64(p.Sprite.Bounds().Dx())-5, YCoordinate: p.YCoordinate}
}

func (p *Player) deleteLaser(windowW float64, windowH float64){
	if p.LeftLaser != nil && p.RightLaser != nil{
		if p.LeftLaser.isLaserOutOfBounds(windowW, windowH) && p.RightLaser.isLaserOutOfBounds(windowW, windowH){
			p.LeftLaser=nil
			p.RightLaser=nil
		}
	}
}