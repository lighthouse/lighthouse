import { Component } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { catchError, retry } from 'rxjs/operators';
import { Observable, throwError } from 'rxjs';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { YamlModalComponent } from './yaml-modal/yaml-modal.component'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  chartName: string = ''
  loading: boolean = false
  error: boolean = false
  objects: KubernetesObject[] = []
  yaml: string = ""

  constructor(private http: HttpClient, private modalService: NgbModal) { }

  showYAML(rawString: string) {
    const modalRef = this.modalService.open(YamlModalComponent, { size: 'lg' })
    modalRef.componentInstance.yaml = rawString
  }

  onVisualize() {
    this.loading = true
    this.error = false

    var chart = {
      name: this.chartName.trim()
    }

    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json'
      })
    }

    this.http.post<KubernetesObject[]>('/chart', chart, httpOptions).subscribe(
      (response) => {
        this.objects = response
        this.loading = false
      },
      (err) => {
        this.error = true
        this.loading = false
      }
    )
  }
}

class KubernetesObject {
  kind: string;
  rawString: string;
  metadata: Metadata;
  yamlState: Boolean;

  public showYAML() {
    alert(this.rawString)
  }
}

class Metadata {
  name: string;
}
