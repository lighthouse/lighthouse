import { Component, OnInit, Input } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-yaml-modal',
  templateUrl: './yaml-modal.component.html',
  styleUrls: ['./yaml-modal.component.css']
})
export class YamlModalComponent implements OnInit {

  @Input()yaml: string;

  constructor(
    public activeModal: NgbActiveModal
  ) { }

  closeModal() {
    this.activeModal.close('Modal Closed');
  }
  
  ngOnInit() {
  }

}
