import { Component } from '@angular/core';
import { HeaderService } from '../../core/services/header.service';

@Component({
  selector: 'app-header',
  imports: [],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss'
})
export class HeaderComponent {
  classHeader = "sticky-top";
  constructor(private headerService: HeaderService){}

  ngOnInit(){

    this.headerService.isHidden$.subscribe((hidden) => {
      if (hidden) {
        this.classHeader = "";
        console.log(this.classHeader);
      } else {
        this.classHeader = "sticky-top";
      }
    });
  }

}
