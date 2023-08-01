using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Enemy : MonoBehaviour
{
    // Start is called before the first frame update

    public float speed = 5;
    Vector3 dir;

    // 폭발 공장 주소
    public GameObject explosionFactory;

    void Start()
    {

        
    }
    void OnEnable()
    {

        int randValue = Random.Range(0, 10);
        if (randValue < 3)
        {
            // 플레이어 방향
            GameObject target = GameObject.Find("Player");
            dir = target.transform.position - transform.position;
            dir.Normalize();
        }
        else
        {
            // 아래 방향
            dir = Vector3.down;
        }
    }
    // Update is called once per frame
    void Update()
    {

        transform.position += dir * speed * Time.deltaTime;
    }

    // 충돌 판정 시
    private void OnCollisionEnter(Collision other)
    {

        // 폭발 이펙트를 발생시킴
        GameObject explosion = Instantiate(explosionFactory);

        // 폭발 효과를 위치시킴
        explosion.transform.position = transform.position;


        // 부딪히면 비활성화시켜서 다시 탄창에 넣는다

        // 부딪힌 상대의 이름이 Bullet이라면
        if (other.gameObject.name.Contains("Bullet"))
        {
            // 부딪힌 물체 비활성화
            other.gameObject.SetActive(false);
        }

        // 아니라면 제거
        else
        {
            Destroy(other.gameObject);
        }

        // Destroy(gameObject);
        gameObject.SetActive(false); // 풀에 자원 반납

        PlayerFire player = GameObject.Find("Player").GetComponent<PlayerFire>();
        player.bulletObjectPool.Add(other.gameObject);

        // 적을 잡을 때마다 점수 표시

        ScoreManager.Instance.Score++;

    }



}
